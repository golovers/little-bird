package api

import (
	"encoding/gob"
	"errors"
	"net/http"
	"net/url"

	plus "google.golang.org/api/plus/v1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	"os"

	uuid "github.com/satori/go.uuid"
)

const (
	defaultSessionID        = "default"
	googleProfileSessionKey = "google_profile"
	oauthTokenSessionKey    = "oauth_token"
	oauthFlowRedirectKey    = "redirect"
)

func init() {
	gob.Register(&oauth2.Token{})
	gob.Register(&Profile{})
}

// loginHandler initiates an OAuth flow to authenticate the user.
func loginHandler(w http.ResponseWriter, r *http.Request) *appError {
	var err error
	sessionID := uuid.Must(uuid.NewV4(), err).String()
	if err != nil {
		return appErrorf(err, "could not create oauth session: %v", err)
	}

	session, err := sessionStore.New(r, sessionID)
	if err != nil {
		return appErrorf(err, "could not create oauth session: %v", err)
	}
	session.Options.MaxAge = 10 * 60 // 10 minutes

	redirectURL, err := validateRedirectURL(r.FormValue("redirect"))
	if err != nil {
		return appErrorf(err, "invalid redirect URL: %v", err)
	}
	session.Values[oauthFlowRedirectKey] = redirectURL

	if err := session.Save(r, w); err != nil {
		return appErrorf(err, "could not save session: %v", err)
	}

	// Use the session ID for the "state" parameter.
	// This protects against CSRF (cross-site request forgery).
	// See https://godoc.org/golang.org/x/oauth2#Config.AuthCodeURL for more detail.
	url := oauthConfig.AuthCodeURL(sessionID, oauth2.ApprovalForce,
		oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusFound)
	http.SetCookie(w, http.Cookie{
		Name:  "authorized",
		Value: "true",
	})
	return nil
}

// validateRedirectURL checks that the URL provided is valid.
// If the URL is missing, redirect the user to the application's root.
// The URL must not be absolute (i.e., the URL must refer to a path within this
// application).
func validateRedirectURL(path string) (string, error) {
	if path == "" {
		return "/", nil
	}
	// ensure redirect URL is valid and not pointing to a different server.
	parsedURL, err := url.Parse(path)
	if err != nil {
		return "/", err
	}
	if parsedURL.IsAbs() {
		return "/", errors.New("URL must not be absolute")
	}
	return path, nil
}

// oauthCallbackHandler completes the OAuth flow, retreives the user's profile
// information and stores it in a session.
func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) *appError {
	oauthFlowSession, err := sessionStore.Get(r, r.FormValue("state"))
	if err != nil {
		return appErrorf(err, "invalid state parameter. try logging in again.")
	}

	redirectURL, ok := oauthFlowSession.Values[oauthFlowRedirectKey].(string)
	// validate this callback request came from the app.
	if !ok {
		return appErrorf(err, "invalid state parameter. try logging in again.")
	}

	code := r.FormValue("code")
	tok, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return appErrorf(err, "could not get auth token: %v", err)
	}

	session, err := sessionStore.New(r, defaultSessionID)
	if err != nil {
		return appErrorf(err, "could not get default session: %v", err)
	}

	ctx := context.Background()
	profile, err := fetchProfile(ctx, tok)
	if err != nil {
		return appErrorf(err, "could not fetch Google profile: %v", err)
	}

	session.Values[oauthTokenSessionKey] = tok
	session.Values[googleProfileSessionKey] = stripProfile(profile)
	if err := session.Save(r, w); err != nil {
		return appErrorf(err, "could not save session: %v", err)
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
	return nil
}

// fetchProfile retrieves the Google+ profile of the user associated with the
// provided OAuth token.
func fetchProfile(ctx context.Context, tok *oauth2.Token) (*plus.Person, error) {
	client := oauth2.NewClient(ctx, oauthConfig.TokenSource(ctx, tok))
	srv, err := plus.New(client)
	if err != nil {
		return nil, err
	}
	return srv.People.Get("me").Do()
}

// logoutHandler clears the default session.
func logoutHandler(w http.ResponseWriter, r *http.Request) *appError {
	session, err := sessionStore.New(r, defaultSessionID)
	if err != nil {
		return appErrorf(err, "could not get default session: %v", err)
	}
	session.Options.MaxAge = -1 // clear session.
	if err := session.Save(r, w); err != nil {
		return appErrorf(err, "could not save session: %v", err)
	}
	redirectURL := r.FormValue("redirect")
	if redirectURL == "" {
		redirectURL = "/"
	}
	http.Redirect(w, r, redirectURL, http.StatusFound)
	http.SetCookie(w, http.Cookie{
		Name:  "authorized",
		Value: "false",
	})
	return nil
}

// profileFromSession retreives the Google+ profile from the default session.
// Returns nil if the profile cannot be retreived (e.g. user is logged out).
func profileFromSession(r *http.Request) *Profile {
	// this should only be used for testing when have no internet connection
	if os.Getenv("LITTLE_BIRD_IGNORE_AUTH") == "true" {
		return &Profile{
			ID:          "anonymous",
			DisplayName: "Anonymous",
		}
	}
	session, err := sessionStore.Get(r, defaultSessionID)
	if err != nil {
		return nil
	}
	tok, ok := session.Values[oauthTokenSessionKey].(*oauth2.Token)
	if !ok || !tok.Valid() {
		return nil
	}
	profile, ok := session.Values[googleProfileSessionKey].(*Profile)
	if !ok {
		return nil
	}
	return profile
}

//Profile hold user profile information
type Profile struct {
	ID, DisplayName, ImageURL string
}

// stripProfile returns a subset of a plus.Person.
func stripProfile(p *plus.Person) *Profile {
	return &Profile{
		ID:          p.Id,
		DisplayName: p.DisplayName,
		ImageURL:    p.Image.Url,
	}
}

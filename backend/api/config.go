package api

import (
	"gitlab.com/7chip/little-bird/backend/common"
	"gopkg.in/mgo.v2"

	"github.com/gorilla/sessions"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig  *oauth2.Config
	sessionStore sessions.Store
	_            mgo.Session
	cfg          Cfg
)

//Cfg wrapper of configurations
type Cfg struct {
	HTTPAddress        string `envconfig:"HTTP_ADDRESS"`
	OAuth2ClientID     string `envconfig:"OAUTH2_CLIENT_ID"`
	OAuth2ClientSecret string `envconfig:"OAUTH2_CLIENT_SECRET"`
	OAuth2Callback     string `envconfig:"OAUTH2_CALLBACK"`
}

func init() {
	common.LoadCfgFromEnv(&cfg)

	oauthConfig = configureOAuthClient(cfg.OAuth2ClientID, cfg.OAuth2ClientSecret)
	cookieStore := sessions.NewCookieStore([]byte("Little Bird - August 20, 2018 - Spread out the world with little things"))
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	sessionStore = cookieStore
}

// configureOAuthClient https://developers.google.com/identity/sign-in/web/sign-in
func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := cfg.OAuth2Callback
	if redirectURL == "" {
		redirectURL = "http://localhost:8080/oauth2callback"
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

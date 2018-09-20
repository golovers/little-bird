package api

import (
	"github.com/gorilla/sessions"
	"gitlab.com/koffee/little-bird/backend/core"

	"gopkg.in/mgo.v2"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfig  *oauth2.Config
	sessionStore sessions.Store
	_            mgo.Session
	conf         Conf
	gw           GW
)

//Conf wrapper of configurations
type Conf struct {
	OAuth2ClientID     string `envconfig:"OAUTH2_CLIENT_ID"`
	OAuth2ClientSecret string `envconfig:"OAUTH2_CLIENT_SECRET"`
	OAuth2Callback     string `envconfig:"OAUTH2_CALLBACK"`
}

func init() {
	core.LoadEnvConfig(&conf)

	oauthConfig = configureOAuthClient(conf.OAuth2ClientID, conf.OAuth2ClientSecret)
	cookieStore := sessions.NewCookieStore([]byte("Little Bird - August 20, 2018 - Spread out the world with little things"))
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	sessionStore = cookieStore

	var err error
	gw, err = NewGWService()
	if err != nil {
		panic("failed to create gw: " + err.Error())
	}
}

// configureOAuthClient https://developers.google.com/identity/sign-in/web/sign-in
func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := conf.OAuth2Callback
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

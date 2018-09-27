package handlers

import (
	"log"
	"os"

	"github.com/golovers/little-bird/backend/core"
	"github.com/gorilla/sessions"
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
	OAuth2Callback     string `envconfig:"OAUTH2_CALLBACK" default:"http://localhost:8080/oauth2callback"`
	SessionStorePath   string `envconfig:"SESSIONS_STORE_PATH" default:"_db/sessions/"`
}

func init() {
	core.LoadEnvConfig(&conf)

	oauthConfig = configureOAuthClient(conf.OAuth2ClientID, conf.OAuth2ClientSecret)
	sessionStore = configureSessionStore(conf.SessionStorePath)

	var err error
	gw, err = NewGW()
	if err != nil {
		panic("failed to create gw: " + err.Error())
	}
}

func configureSessionStore(path string) sessions.Store {
	k := []byte("Little Bird - August 20, 2018 - Spread out the little things ^_^") //64 bytes
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println("auth config: failed to create sessions file system store, fallback to use cookie store")
		s := sessions.NewCookieStore(k)
		s.Options = &sessions.Options{
			HttpOnly: true,
		}
		return s
	}
	s := sessions.NewFilesystemStore(path, k)
	s.Options = &sessions.Options{
		HttpOnly: true,
	}
	return s
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

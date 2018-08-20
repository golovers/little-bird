package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// RegisterHandlers register all necessary handlers
func RegisterHandlers() {
	r := mux.NewRouter()

	r.Methods("GET").Path("/").Handler(appHandler(index))
	r.Methods("GET").Path("/post").Handler(appHandler(post))

	r.Methods("GET").Path("/login").Handler(appHandler(loginHandler))
	r.Methods("POST").Path("/logout").Handler(appHandler(logoutHandler))
	r.Methods("GET").Path("/oauth2callback").Handler(appHandler(oauthCallbackHandler))

	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}

// index display the  index page
func index(w http.ResponseWriter, r *http.Request) *appError {
	return indexTmpl.Execute(w, r, "")
}

// post display the  post page
func post(w http.ResponseWriter, r *http.Request) *appError {
	return postTmpl.Execute(w, r, "")
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)

		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}

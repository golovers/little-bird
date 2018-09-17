package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.com/koffee/little-bird/backend/core"
)

// RegisterHandlers register all necessary handlers
func RegisterHandlers() {
	r := mux.NewRouter()
	// REST API
	r.Path("/api/v1/articles").Methods("POST").Handler(appHandler(createArticle))

	// Form request
	r.Methods("GET").Path("/").Handler(appHandler(index))

	r.Path("/articles").Methods("GET").Handler(appHandler(index))
	r.Path("/articles/trending").Methods("GET").Handler(appHandler(trendingArticlesHandler))
	r.Path("/articles/mine").Methods("GET").Handler(appHandler(myArticlesHandler))
	r.Path("/articles/add").Methods("GET").Handler(appHandler(newArticleHandler))
	r.Path("/articles/details/{id:[a-z0-9]+}").Methods("GET").Handler(appHandler(articleDetailsHandler))

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
	articles, err := gw.ListArticle(context.Background(), core.Pagination{})
	if err != nil {
		return appErrorf(err, "failed to list all articles")
	}
	// don't need the markdown details in this case
	for _, a := range articles {
		a.Markdown = ""
	}
	return indexTmpl.Execute(w, r, articles)
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

func appOK(id string, w http.ResponseWriter) *appError {
	data := struct {
		ID string
	}{
		ID: id,
	}
	b, _ := json.Marshal(data)
	w.Write(b)
	return nil
}

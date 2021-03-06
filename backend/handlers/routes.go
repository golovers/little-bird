package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/golovers/little-bird/backend/core"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Register register all necessary handlers
func Register() {
	r := mux.NewRouter()
	// REST API
	r.Path("/api/v1/articles").Methods("POST").Handler(appHandler(createArticle))
	r.Path("/api/v1/articles/{id}/vote").Methods("POST").Handler(appHandler(upVote))
	r.Path("/api/v1/articles/{id}").Methods("PUT").Handler(appHandler(updateArticle))
	r.Path("/api/v1/articles/{id}").Methods("DELETE").Handler(appHandler(deleteArticle))
	r.Path("/api/v1/articles/{id}/comments").Methods("GET").Handler(appHandler(listCommentByArticle))

	r.Path("/api/v1/comments").Methods("POST").Handler(appHandler(createComment))
	r.Path("/api/v1/comments/{id}").Methods("PUT").Handler(appHandler(updateComment))
	r.Path("/api/v1/comments/{id}").Methods("DELETE").Handler(appHandler(deleteComment))

	// Form request
	r.Methods("GET").Path("/").Handler(appHandler(index))

	r.Path("/articles").Methods("GET").Handler(appHandler(index))
	r.Path("/articles/trending").Methods("GET").Handler(appHandler(handleArticleTrending))
	r.Path("/articles/mine").Methods("GET").Handler(appHandler(handleArticleMine))
	r.Path("/articles/add").Methods("GET").Handler(appHandler(handleArticleNew))
	r.Path("/articles/{id:[a-z0-9]+}/details").Methods("GET").Handler(appHandler(handleArticleDetail))
	r.Path("/articles/{id:[a-z0-9]+}/edit").Methods("GET").Handler(appHandler(handleArticleEdit))

	r.Methods("GET").Path("/login").Handler(appHandler(loginHandler))
	r.Methods("POST").Path("/logout").Handler(appHandler(logoutHandler))
	r.Methods("GET").Path("/oauth2callback").Handler(appHandler(oauthCallbackHandler))

	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	templatesStatic := staticFilterHandler(http.FileServer(http.Dir("./templates/")), []string{".css", ".js", ".png", ".jpg"})
	r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", templatesStatic))

	r.NotFoundHandler = appHandler(handleNotFound)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))
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

func handleNotFound(w http.ResponseWriter, r *http.Request) *appError {
	return notFoundTmpl.Execute(w, r, "")
}

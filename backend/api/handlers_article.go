package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func articleDetailsHandler(w http.ResponseWriter, r *http.Request) *appError {
	vars := mux.Vars(r)
	id := vars["id"]
	p, err := gw.GetArticle(context.Background(), id)
	if err != nil {
		return appErrorf(err, "failed to load article %s", id)
	}
	return articleDetailsTmpl.Execute(w, r, p)
}

func newArticleHandler(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		http.Redirect(w, r, "/login?redirect=/articles/add", http.StatusFound)
		return nil
	}
	return newArticleTmpl.Execute(w, r, "")
}

func myArticlesHandler(w http.ResponseWriter, r *http.Request) *appError {
	return myArticlesTmpl.Execute(w, r, "")
}

func trendingArticlesHandler(w http.ResponseWriter, r *http.Request) *appError {
	return trendingArticlesTmpl.Execute(w, r, "")
}

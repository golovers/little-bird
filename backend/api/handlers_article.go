package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/koffee/little-bird/backend/core"
	"golang.org/x/net/context"
)

func handleArticleDetail(w http.ResponseWriter, r *http.Request) *appError {
	vars := mux.Vars(r)
	id := vars["id"]
	p, err := gw.GetArticle(context.Background(), id)
	if err != nil {
		return appErrorf(err, "failed to load article %s", id)
	}
	return articleDetailsTmpl.Execute(w, r, p)
}

func handleArticleNew(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		http.Redirect(w, r, "/login?redirect=/articles/add", http.StatusFound)
		return nil
	}
	return newArticleTmpl.Execute(w, r, "")
}

func handleArticleMine(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		http.Redirect(w, r, "/login?redirect=/articles/mine", http.StatusFound)
		return nil
	}
	articles, err := gw.ListArticleCreatedBy(context.Background(), profile.ID)
	if err != nil {
		return appErrorf(err, "failed to list my articles")
	}
	return myArticlesTmpl.Execute(w, r, articles)
}

func handleArticleTrending(w http.ResponseWriter, r *http.Request) *appError {
	articles, err := gw.TrendingArticle(context.Background(), core.Pagination{})
	if err != nil {
		return appErrorf(err, "failed to list all articles")
	}
	for _, a := range articles {
		a.Markdown = ""
	}
	return trendingArticlesTmpl.Execute(w, r, articles)
}

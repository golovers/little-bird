package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"gitlab.com/koffee/little-bird/backend/core"
)

func createArticle(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return appErrorf(fmt.Errorf("unauthorized"), "unauthorized")
	}
	var article core.Article
	if err := decode(r.Body, &article); err != nil {
		return appErrorf(err, "invalid input")
	}
	article.LastUpdate = time.Now()
	article.CreatedByID = profile.ID
	article.CreatedBy = profile.DisplayName

	id, err := gw.CreateArticle(context.Background(), &article)
	if err != nil {
		return appErrorf(err, "internal server error")
	}
	responseWithData(w, http.StatusOK, map[string]string{
		"ID": id,
	})
	return nil
}

func upVote(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		http.Redirect(w, r, "/login?redirect="+r.URL.Path, http.StatusFound)
	}
	v := &core.Vote{}
	v.ArticleID = mux.Vars(r)["id"]
	v.CreatedBy = profile.DisplayName
	v.CreatedByID = profile.ID
	v.LastUpdate = time.Now()

	_, err := gw.CreateVote(context.Background(), v)
	if err != nil {
		return appErrorf(err, "internal server error")
	}
	article, _ := gw.GetArticle(context.Background(), v.ArticleID)
	responseWithData(w, http.StatusOK, map[string]interface{}{"count": article.VoteCount})
	return nil
}

func updateArticle(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return appErrorf(fmt.Errorf("unauthorized"), "unauthorized")
	}
	var article core.Article
	if err := decode(r.Body, &article); err != nil {
		return appErrorf(err, "invalid input")
	}
	existingArticle, err := gw.GetArticle(context.Background(), article.ID)
	if err != nil {
		return appErrorf(err, "could not get existing article for update")
	}
	if profile.ID != existingArticle.CreatedByID {
		return appErrorf(fmt.Errorf("unauthorized"), "you are not allowed to edit this article")
	}
	existingArticle.Markdown = article.Markdown
	existingArticle.Content = article.Content
	existingArticle.LastUpdate = time.Now()

	err = gw.UpdateArticle(context.Background(), existingArticle.Article)
	if err != nil {
		return appErrorf(err, "internal server error")
	}
	responseWithData(w, http.StatusOK, map[string]string{
		"ID": existingArticle.ID,
	})
	return nil
}

func deleteArticle(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return appErrorf(fmt.Errorf("unauthorized"), "unauthorized")
	}
	article, err := gw.GetArticle(context.Background(), mux.Vars(r)["id"])
	if err != nil {
		return appErrorf(err, "could not found the given article")
	}
	if profile.ID != article.CreatedByID {
		return appErrorf(fmt.Errorf("unauthorized"), "you are not allowed to delete this article")
	}

	err = gw.DeleteArticle(context.Background(), article.ID)
	//TODO remove all relevants comments, votes,...
	if err != nil {
		return appErrorf(err, "internal server error")
	}
	return nil
}

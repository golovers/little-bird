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
		return appErrorf(fmt.Errorf("unauthorized"), "unauthorized")
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

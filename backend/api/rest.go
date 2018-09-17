package api

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
	"gitlab.com/koffee/little-bird/backend/core"
	"context"
)

func createArticle(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return appErrorf(fmt.Errorf("unauthorized"), "unauthorized")
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return appErrorf(err, "failed to read input body")
	}
	var article *core.Article
	err = json.Unmarshal(b, &article)
	if err != nil {
		return appErrorf(err, "invalid input")
	}
	article.LastUpdate = time.Now()
	article.CreatedByID = profile.ID
	article.CreatedBy = profile.DisplayName
	id, err := gw.CreateArticle (context.Background(), article)
	if err != nil {
		return appErrorf(err, "internal server error")
	}
	w.WriteHeader(http.StatusOK)
	return appOK(id, w)
}


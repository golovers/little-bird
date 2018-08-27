package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"gitlab.com/7chip/little-bird/backend/post"

	"github.com/sirupsen/logrus"
)

func postDetails(w http.ResponseWriter, r *http.Request) *appError {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return appErrorf(err, "invalid input")
	}
	p, err := post.Get(id)
	if err != nil {
		return appErrorf(err, "failed to load post from db")
	}
	return postTmpl.Execute(w, r, p)
}

func postAdd(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		http.Redirect(w, r, "/login?redirect=/posts/add", http.StatusFound)
		return nil
	}
	return postAddTmpl.Execute(w, r, "")
}

func postMine(w http.ResponseWriter, r *http.Request) *appError {
	return postMineTmpl.Execute(w, r, "")
}

func postTrending(w http.ResponseWriter, r *http.Request) *appError {
	return postTrendingTmpl.Execute(w, r, "")
}

func postSave(w http.ResponseWriter, r *http.Request) *appError {
	title := r.FormValue("postTitle")
	content := r.FormValue("postContent")
	logrus.Info(title, content)
	return nil
}

func apiPostAdd(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return appErrorf(fmt.Errorf("unauthorized"), "unauthorized")
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return appErrorf(err, "failed to read input body")
	}
	var p post.Post
	err = json.Unmarshal(b, &p)
	if err != nil {
		return appErrorf(err, "invalid input")
	}
	p.LastUpdate = time.Now()
	p.CreatedByID = profile.ID
	p.CreatedBy = profile.DisplayName
	id, err := post.Add(&p)
	if err != nil {
		return appErrorf(err, "internal server error")
	}
	w.WriteHeader(http.StatusOK)
	return appOK(id, w)
}

type Post struct {
	ID         int64
	Title      string
	Content    string
	LastUpdate time.Time

	CreatedBy   string
	CreatedByID string
}

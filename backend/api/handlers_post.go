package api

import (
	"encoding/json"
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
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return appErrorf(err, "failed to read input body")
	}
	var p post.Post
	err = json.Unmarshal(b, &p)
	if err != nil {
		return appErrorf(err, "invalid input")
	}
	post.Add(&p)
	w.WriteHeader(http.StatusOK)
	buff, err := json.Marshal(p)
	if err != nil {
		return appErrorf(err, "internal server error")
	}
	w.Write(buff)
	return nil
}

type Post struct {
	ID         int64
	Title      string
	Content    string
	LastUpdate time.Time

	CreatedBy   string
	CreatedByID string
}

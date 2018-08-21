package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func post(w http.ResponseWriter, r *http.Request) *appError {
	return postTmpl.Execute(w, r, "")
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

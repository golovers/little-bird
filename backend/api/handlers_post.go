package api

import "net/http"

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

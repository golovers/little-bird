package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"gitlab.com/koffee/little-bird/backend/core"
)

var (
	errUnAuthorized   = appErrorf(fmt.Errorf("unauthorized"), "unauthorized")
	errInternalServer = func(err error) *appError { return appErrorf(err, "internal server error") }
	errInvalidInput   = func(err error) *appError { return appErrorf(err, "invalid input") }
)

func createArticle(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return errUnAuthorized
	}
	var article core.Article
	if err := decode(r.Body, &article); err != nil {
		return errInvalidInput(err)
	}
	article.LastUpdate = time.Now()
	article.CreatedByID = profile.ID
	article.CreatedBy = profile.DisplayName

	id, err := gw.CreateArticle(context.Background(), &article)
	if err != nil {
		return errInternalServer(err)
	}
	responseWithData(w, http.StatusOK, map[string]string{
		"ID": id,
	})
	return nil
}

func upVote(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return errUnAuthorized
	}
	v := &core.Vote{}
	v.ArticleID = mux.Vars(r)["id"]
	v.CreatedBy = profile.DisplayName
	v.CreatedByID = profile.ID
	v.LastUpdate = time.Now()

	_, err := gw.CreateVote(context.Background(), v)
	if err != nil {
		return errInternalServer(err)
	}
	article, _ := gw.GetArticle(context.Background(), v.ArticleID)
	responseWithData(w, http.StatusOK, map[string]interface{}{"count": article.VoteCount})
	return nil
}

func updateArticle(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return errUnAuthorized
	}
	var article core.Article
	if err := decode(r.Body, &article); err != nil {
		return errInvalidInput(err)
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

	err = gw.UpdateArticle(context.Background(), existingArticle)
	if err != nil {
		return errInternalServer(err)
	}
	responseWithData(w, http.StatusOK, map[string]string{
		"ID": existingArticle.ID,
	})
	return nil
}

func deleteArticle(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return errUnAuthorized
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
		return errInternalServer(err)
	}
	return nil
}

func listCommentByArticle(w http.ResponseWriter, r *http.Request) *appError {
	comments, err := gw.ListCommentByArticle(context.Background(), mux.Vars(r)["id"])
	log.Println(mux.Vars(r)["id"])
	if err != nil {
		return appErrorf(err, "failed to get comments")
	}
	responseWithData(w, http.StatusOK, comments)
	return nil
}

func createComment(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return errUnAuthorized
	}
	var c core.Comment
	if err := decode(r.Body, &c); err != nil {
		return errInvalidInput(err)
	}
	if c.ArticleID == "" {
		return appErrorf(errors.New("articleID is missing"), "invalid input")
	}
	c.LastUpdate = time.Now()
	c.Created = time.Now()
	c.CreatedBy = profile.DisplayName
	c.CreatedByID = profile.ID

	id, err := gw.CreateComment(context.Background(), &c)
	if err != nil {
		return errInternalServer(err)
	}
	responseWithData(w, http.StatusOK, map[string]string{
		"ID": id,
	})
	return nil
}

func updateComment(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return errUnAuthorized
	}
	var comment core.Comment
	if err := decode(r.Body, &comment); err != nil {
		return errInvalidInput(err)
	}
	existingComment, err := gw.GetComment(context.Background(), mux.Vars(r)["id"])
	if err != nil {
		return appErrorf(err, "could not get existing comment for update")
	}
	if profile.ID != existingComment.CreatedByID {
		return appErrorf(fmt.Errorf("unauthorized"), "you are not allowed to edit this comment")
	}
	existingComment.Content = comment.Content
	existingComment.LastUpdate = time.Now()

	err = gw.UpdateComment(context.Background(), existingComment)
	if err != nil {
		return errInternalServer(err)
	}
	responseWithData(w, http.StatusOK, map[string]string{
		"ID": existingComment.ID,
	})
	return nil
}

func deleteComment(w http.ResponseWriter, r *http.Request) *appError {
	profile := profileFromSession(r)
	if profile == nil {
		return errUnAuthorized
	}
	comment, err := gw.GetComment(context.Background(), mux.Vars(r)["id"])
	if err != nil {
		return appErrorf(err, "could not found the given comment")
	}
	if profile.ID != comment.CreatedByID {
		return appErrorf(fmt.Errorf("unauthorized"), "you are not allowed to delete this comment")
	}

	err = gw.DeleteComment(context.Background(), comment.ID)
	if err != nil {
		return errInternalServer(err)
	}
	return nil
}

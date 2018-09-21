package articles

import (
	"log"
	"testing"
	"time"

	"gitlab.com/koffee/little-bird/backend/core"
	"golang.org/x/net/context"
)

func TestArticles(t *testing.T) {
	a := &core.Article{
		Title:       "Little Bird Article",
		Content:     "#Little Bird content",
		CreatedBy:   "pthethanh",
		CreatedByID: "pthethanh_id",
		LastUpdate:  time.Now(),
	}

	service, err := NewArticleService()
	if err != nil {
		t.Fatalf("articles: failed to create article service: %v", err)
	}
	id, err := service.Create(context.Background(), a)
	if err != nil {
		t.Fatalf("articles: failed to create new article: %v", err)
	}
	log.Printf("newly created article: %s\n", id)
	newArticle, err := service.Get(context.Background(), id)
	if err != nil {
		t.Fatalf("articles: failed to get article: %v", err)
	}
	log.Printf("found the article: %s, with content: %s\n", newArticle.Title, newArticle.Content)
	articles, err := service.List(context.Background(), core.Pagination{})
	if err != nil {
		t.Fatalf("articles: failed to get list of articles: %v", err)
	}
	log.Printf("found %d articles\n", len(articles))
	newArticles, err := service.ListCreatedBy(context.Background(), a.CreatedByID)
	if err != nil {
		t.Fatalf("articles: failed to list by created by: %v", err)
	}
	log.Printf("list by created by. Found %d articles\n", len(newArticles))
	a.Content = "new content"

	err = service.Update(context.Background(), a)
	if err != nil {
		t.Fatalf("articles: failed to update")
	}
	updatedArticle, _ := service.Get(context.Background(), a.ID)
	if updatedArticle.Content != "new content" {
		t.Fatalf("articles: failed to update the article")
	}
	err = service.Delete(context.Background(), id)
	if err != nil {
		t.Fatalf("articles: failed to delete existing article: %v", err)
	}
}

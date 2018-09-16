package core

import (
	"context"
	"time"
)

// Article prepresent an article
type Article struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Content     string    `json:"content,omitempty"`
	LastUpdate  time.Time `json:"last_update,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	CreatedByID string    `json:"created_by_id,omitempty"`
}

// ArticleServicer is the server API for ArticleService service.
type ArticleServicer interface {
	// List articles
	List(context.Context, Pagination) ([]*Article, error)
	// ListCreatedBy list articles by a specific user
	ListCreatedBy(context.Context, string) ([]*Article, error)
	// Delete delete an article
	Delete(context.Context, string) error
	// Get get a specific article
	Get(context.Context, string) (*Article, error)
	// Update update a specific article
	Update(context.Context, *Article) error
	// Update update a specific article
	Create(context.Context, *Article) (string, error)
}

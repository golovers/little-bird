package core

import (
	"context"
	"time"
)

// Comment represent content of a comment
type Comment struct {
	ID          string    `json:"id,omitempty"`
	ArticleID   string    `json:"article_id,omitempty"`
	Content     string    `json:"content,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	CreatedByID string    `json:"created_by_id,omitempty"`
	LastUpdate  time.Time `json:"last_update,omitempty"`
}

// CommentServicer provide operations on comments
type CommentServicer interface {
	ListByArticle(ctx context.Context, articleID string) ([]*Comment, error)
	CountByArticles(ctx context.Context, articleIDs ...string) (map[string]int, error)
	Update(ctx context.Context, c *Comment) error
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, c *Comment) (id string, err error)
}

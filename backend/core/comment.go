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
	CreatedBy   string    `json:"fullname,omitempty"`
	CreatedByID string    `json:"creator,omitempty"`
	LastUpdate  time.Time `json:"modified,omitempty"`
	Parent      string    `json:"parent,omitempty"`
	Created     time.Time `json:"created,omitempty"`

	ProfileURL        string `json:"profile_url,omitempty"`
	ProfilePictureURL string `json:"profile_picture_url,omitempty"`
	UpvoteCount       int    `json:"upvote_count,omitempty"`

	// temporary leave them here...they might be moved to handlers level
	CreatedByCurrentUser bool `json:"created_by_current_user,omitempty"`
	UserHasUpvoted       bool `json:"user_has_upvoted,omitempty"`
}

// CommentServicer provide operations on comments
type CommentServicer interface {
	ListByArticle(ctx context.Context, articleID string) ([]*Comment, error)
	CountByArticles(ctx context.Context, articleIDs ...string) (map[string]int, error)
	Update(ctx context.Context, c *Comment) error
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, c *Comment) (id string, err error)
	Get(ctx context.Context, id string) (*Comment, error)
}

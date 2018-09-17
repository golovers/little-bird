package core

import "context"

// Vote represent content of a vote for a specific article
type Vote struct {
	ID          string `json:"id,omitempty"`
	ArticleID   string `json:"article_id,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	CreatedByID string `json:"created_by_id,omitempty"`
}

//VoteServicer provide operations on voting
type VoteServicer interface {
	Create(ctx context.Context, v *Vote) (id string, err error)
	CountByArticles(ctx context.Context, articleID ...string) (map[string]int, error)
}

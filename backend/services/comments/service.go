package comments

import (
	"context"

	"github.com/golovers/little-bird/backend/core"
)

// Repository database access to comments
type Repository interface {
	ListByArticle(articleID string) ([]*core.Comment, error)
	Create(b *core.Comment) (id string, err error)
	Delete(id string) error
	Get(id string) (*core.Comment, error)
	Update(b *core.Comment) error
	Close()
}

type commentService struct {
	repo Repository
}

//NewCommentService create default implementation of comment service
func NewCommentService() (core.CommentServicer, error) {
	db, err := newMongoDB()
	if err != nil {
		return &commentService{}, err
	}
	return &commentService{
		repo: db,
	}, nil
}

func (s *commentService) ListByArticle(ctx context.Context, articleID string) ([]*core.Comment, error) {
	return s.repo.ListByArticle(articleID)
}

func (s *commentService) CountByArticles(ctx context.Context, articleIDs ...string) (map[string]int, error) {
	//TODO improve me
	rs := make(map[string]int)
	for _, id := range articleIDs {
		cmts, err := s.ListByArticle(ctx, id)
		if err != nil {
			rs[id] = 0
		}
		rs[id] = len(cmts)
	}
	return rs, nil
}

func (s *commentService) Update(ctx context.Context, c *core.Comment) error {
	return s.repo.Update(c)
}

func (s *commentService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}

func (s *commentService) Create(ctx context.Context, c *core.Comment) (id string, err error) {
	return s.repo.Create(c)
}

func (s *commentService) Get(ctx context.Context, id string) (*core.Comment, error) {
	return s.repo.Get(id)
}

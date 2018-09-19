package votes

import (
	"context"

	"gitlab.com/koffee/little-bird/backend/core"
)

var _ core.VoteServicer = (*voteService)(nil)

// Repository database access to comments
type Repository interface {
	ListByArticle(articleID string) ([]*core.Vote, error)
	Create(b *core.Vote) (id string, err error)
	Get(id string) (*core.Vote, error)
	GetByArticleAndUserID(articleID string, userID string) (*core.Vote, error)
	Close()
}

type voteService struct {
	repo Repository
}

//NewVoteService return default implementation of vote service
func NewVoteService() (core.VoteServicer, error) {
	db, err := newMongoDB()
	if err != nil {
		return &voteService{}, err
	}
	return &voteService{
		repo: db,
	}, nil
}

func (s *voteService) Create(ctx context.Context, v *core.Vote) (string, error) {
	if v, _ := s.repo.GetByArticleAndUserID(v.ArticleID, v.CreatedByID); v.ID != "" {
		return v.ID, nil
	}
	return s.repo.Create(v)
}

func (s *voteService) CountByArticles(ctx context.Context, articleID ...string) (map[string]int, error) {
	rs := make(map[string]int)
	for _, id := range articleID {
		if vs, err := s.repo.ListByArticle(id); err == nil {
			rs[id] = len(vs)
		}
	}
	return rs, nil
}

package articles

import (
	"context"

	"gitlab.com/koffee/little-bird/backend/core"
)

// Repository database access to articles
type Repository interface {
	// List returns a list of posts, ordered by title.
	List(offset, limit int64) ([]*core.Article, error)

	// ListByCreatedBy returns a list of posts, ordered by title, filtered by
	// the user who created the post entry.
	ListByCreatedBy(userID string) ([]*core.Article, error)

	// Get retrieves a post by its ID.
	Get(id string) (*core.Article, error)

	// Create saves a given post, assigning it a new ID.
	Create(b *core.Article) (id string, err error)

	// Delete removes a given post by its ID.
	Delete(id string) error

	// Update updates the entry for a given post.
	Update(b *core.Article) error

	// Close closes the database, freeing up any available resources.
	Close()
}

// make sure article service implement article service server
var _ core.ArticleServicer = &articleService{}

type articleService struct {
	repo Repository
}

// NewArticleService create new ready-to-use article service with default underlying database (mongodb)
func NewArticleService() (core.ArticleServicer, error) {
	repo, err := newMongoDB()
	return &articleService{
		repo: repo,
	}, err
}

// List articles
func (s *articleService) List(ctx context.Context, p core.Pagination) ([]*core.Article, error) {
	articles, err := s.repo.List(p.Offset, p.Limit)
	if err != nil {
		return []*core.Article{}, err
	}
	return articles, nil
}

// ListCreaterepoy list articles by a specific user
func (s *articleService) ListCreatedBy(ctx context.Context, ownerID string) ([]*core.Article, error) {
	articles, err := s.repo.ListByCreatedBy(ownerID)
	if err != nil {
		return []*core.Article{}, err
	}
	return articles, nil
}

// Delete delete an article
func (s *articleService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(id)
}

// Get get a specific article
func (s *articleService) Get(ctx context.Context, id string) (*core.Article, error) {
	article, err := s.repo.Get(id)
	if err != nil {
		return &core.Article{}, err
	}
	return article, nil
}

// Update update a specific article
func (s *articleService) Update(ctx context.Context, a *core.Article) error {
	return s.repo.Update(a)
}

func (s *articleService) Create(ctx context.Context, a *core.Article) (string, error) {
	return s.repo.Create(a)
}

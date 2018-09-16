package articles

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "gitlab.com/koffee/little-bird/backend/apis-go/article/v1"
)

// Repository database access to articles
type Repository interface {
	// List returns a list of posts, ordered by title.
	List(offset, limit int64) ([]*pb.Article, error)

	// ListByCreaterepoy returns a list of posts, ordered by title, filtered by
	// the user who created the post entry.
	ListByCreaterepoy(userID string) ([]*pb.Article, error)

	// Get retrieves a post by its ID.
	Get(id string) (*pb.Article, error)

	// Create saves a given post, assigning it a new ID.
	Create(b *pb.Article) (id string, err error)

	// Delete removes a given post by its ID.
	Delete(id string) error

	// Update updates the entry for a given post.
	Update(b *pb.Article) error

	// Close closes the database, freeing up any available resources.
	Close()
}

// make sure article service implement article service server
var _ pb.ArticleServiceServer = &articleService{}

type articleService struct {
	repo Repository
}

// List articles
func (s *articleService) List(ctx context.Context, r *pb.ListRequest) (*pb.ListResponse, error) {
	articles, err := s.repo.List(r.Offset, r.Limit)
	if err != nil {
		return &pb.ListResponse{}, err
	}
	return &pb.ListResponse{
		Articles: articles,
	}, nil
}

// ListCreaterepoy list articles by a specific user
func (s *articleService) ListCreatedBy(ctx context.Context, r *pb.ListCreatedByRequest) (*pb.ListResponse, error) {
	articles, err := s.repo.ListByCreaterepoy(r.OwnerId)
	if err != nil {
		return &pb.ListResponse{}, err
	}
	return &pb.ListResponse{
		Articles: articles,
	}, nil
}

// Delete delete an article
func (s *articleService) Delete(ctx context.Context, r *pb.DeleteRequest) (*empty.Empty, error) {
	err := s.repo.Delete(r.GetId())
	return &empty.Empty{}, err
}

// Get get a specific article
func (s *articleService) Get(ctx context.Context, r *pb.GetRequest) (*pb.Article, error) {
	article, err := s.repo.Get(r.Id)
	if err != nil {
		return &pb.Article{}, err
	}
	return article, nil
}

// Update update a specific article
func (s *articleService) Update(ctx context.Context, r *pb.UpdateRequest) (*empty.Empty, error) {
	err := s.repo.Update(r.Article)
	return &empty.Empty{}, err
}

func (s *articleService) Create(ctx context.Context, r *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := s.repo.Create(r.Article)
	return &pb.CreateResponse{
		Id: id,
	}, err
}

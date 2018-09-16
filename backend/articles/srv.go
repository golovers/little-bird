package articles

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "gitlab.com/koffee/little-bird/backend/apis-go/article/v1"
)

// ArticleService serve operations on article
type ArticleService interface {
	// List articles
	ListArticles(context.Context, *pb.ListArticleRequest) (*pb.ListArticleResponse, error)
	// ListArticleCreatedBy list articles by a specific user
	ListArticleCreatedBy(context.Context, *pb.ListArticleCreatedByRequest) (*pb.ListArticleResponse, error)
	// DeleteArticle delete an article
	DeleteArticle(context.Context, *pb.DeleteArticletRequest) (*empty.Empty, error)
	// GetArticle get a specific article
	GetArticle(context.Context, *pb.GetArticletRequest) (*pb.Article, error)
	// UpdateArticle update a specific article
	UpdateArticle(context.Context, *pb.UpdateArticletRequest) (*empty.Empty, error)
	// Trending return list of trending articles
	Trending(context.Context, *pb.TrendingRequest) (*pb.ListArticleResponse, error)
}

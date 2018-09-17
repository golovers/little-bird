package api

import (
	"gitlab.com/koffee/little-bird/backend/articles"
	"gitlab.com/koffee/little-bird/backend/core"
	"gitlab.com/koffee/little-bird/comments"
	"gitlab.com/koffee/little-bird/votes"
	"golang.org/x/net/context"
)

//ArticleOverview hold article information and its statistic data
type ArticleOverview struct {
	*core.Article
}

//ArticleDetails hold details information of an article
type ArticleDetails struct {
	*ArticleOverview
	Comments []*core.Comment
}

//GW join internal services and provide public operations to web or rest api
type GW interface {
	ListArticle(ctx context.Context, p core.Pagination) ([]*ArticleOverview, error)
	ListArticleCreatedBy(ctx context.Context, userID string) ([]*ArticleOverview, error)
	TrendingArticle(ctx context.Context, p core.Pagination) ([]*ArticleOverview, error)
	CreateArticle(ctx context.Context, a *core.Article) (string, error)
	UpdateArticle(ctx context.Context, a *core.Article) error
	DeleteArticle(ctx context.Context, id string) error
	GetArticle(ctx context.Context, id string) (*ArticleDetails, error)

	CreateComment(ctx context.Context, c *core.Comment) (string, error)
	DeleteComment(ctx context.Context, id string) error

	CreateVote(ctx context.Context, v *core.Vote) (string, error)
}

var _ GW = &gwService{}

type gwService struct {
	articleService core.ArticleServicer
	commentService core.CommentServicer
	voteService    core.VoteServicer
}

//NewGWService return a default implementation of GW service
func NewGWService() (GW, error) {
	article, err := articles.NewArticleService()
	if err != nil {
		return &gwService{}, err
	}
	comment, err := comments.NewCommentService()
	if err != nil {
		return &gwService{}, err
	}
	vote, err := votes.NewVoteService()
	if err != nil {
		return &gwService{}, err
	}
	return &gwService{
		articleService: article,
		commentService: comment,
		voteService:    vote,
	}, nil
}

func (gw *gwService) ListArticle(ctx context.Context, p core.Pagination) ([]*ArticleOverview, error) {
	articles, err := gw.articleService.List(ctx, p)
	if err != nil {
		return []*ArticleOverview{}, err
	}
	articlesOverview := make([]*ArticleOverview, len(articles))
	ids := make([]string, 0)
	for _, a := range articles {
		articlesOverview = append(articlesOverview, &ArticleOverview{
			Article: a,
		})
		ids = append(ids, a.ID)
	}
	return articlesOverview, nil
}

func (gw *gwService) ListArticleCreatedBy(ctx context.Context, userID string) ([]*ArticleOverview, error) {
	return []*ArticleOverview{}, nil
}

func (gw *gwService) TrendingArticle(ctx context.Context, p core.Pagination) ([]*ArticleOverview, error) {
	//TODO implement me
	return []*ArticleOverview{}, nil
}

func (gw *gwService) CreateArticle(ctx context.Context, a *core.Article) (string, error) {
	return gw.articleService.Create(ctx, a)
}

func (gw *gwService) UpdateArticle(ctx context.Context, a *core.Article) error {
	return gw.articleService.Update(ctx, a)
}

func (gw *gwService) DeleteArticle(ctx context.Context, id string) error {
	return gw.articleService.Delete(ctx, id)
}

func (gw *gwService) GetArticle(ctx context.Context, id string) (*ArticleDetails, error) {
	//TODO implement me
	return &ArticleDetails{}, nil
}

func (gw *gwService) CreateComment(ctx context.Context, c *core.Comment) (string, error) {
	//TODO update comment count in article
	return gw.commentService.Create(ctx, c)
}

func (gw *gwService) DeleteComment(ctx context.Context, id string) error {
	// TODO update comment count in article
	return gw.commentService.Delete(ctx, id)
}

func (gw *gwService) CreateVote(ctx context.Context, v *core.Vote) (string, error) {
	// TODO update vote count in article
	return gw.voteService.Create(ctx, v)
}
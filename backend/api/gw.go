package api

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/koffee/little-bird/backend/articles"
	"gitlab.com/koffee/little-bird/backend/core"
	"golang.org/x/net/context"
)

//ArticleOverview hold article information and its statistic data
type ArticleOverview struct {
	*core.Article
	*ArticleStatistic
}

//ArticleStatistic hold statistic data of article
type ArticleStatistic struct {
	ViewCount    int      `json:"view_count,omitempty"`
	CommentCount int      `json:"comment_count,omitempty"`
	VoteCount    int      `json:"vote_count,omitempty"`
	Tags         []string `json:"tags,omitempty"`
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
	return &gwService{
		articleService: article,
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
	// update comment count
	comments, err := gw.commentService.CountByArticles(ctx, ids...)
	if err != nil {
		logrus.Errorf("failed to fetch comments: %v", err)
	} else {
		for _, a := range articlesOverview {
			a.CommentCount = comments[a.ID]
		}
	}

	// update vote count
	votes, err := gw.voteService.CountByArticles(ctx, ids...)
	if err != nil {
		logrus.Errorf("failed to fetch votes: %v", err)
	} else {
		for _, a := range articlesOverview {
			a.VoteCount = votes[a.ID]
		}
	}
	// update view count
	// TODO implement me

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
	return gw.commentService.Create(ctx, c)
}

func (gw *gwService) DeleteComment(ctx context.Context, id string) error {
	return gw.commentService.Delete(ctx, id)
}

func (gw *gwService) CreateVote(ctx context.Context, v *core.Vote) (string, error) {
	//TODO implement me
	return gw.voteService.Create(ctx, v)
}

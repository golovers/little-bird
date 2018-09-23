package handlers

import (
	"context"
	"sort"

	"gitlab.com/koffee/little-bird/backend/core"
	"gitlab.com/koffee/little-bird/backend/services/articles"
	"gitlab.com/koffee/little-bird/backend/services/comments"
	"gitlab.com/koffee/little-bird/backend/services/votes"
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

//GW join internal services and provide public operations to web or rest handlers
type GW interface {
	ListArticle(ctx context.Context, p core.Pagination) ([]*ArticleOverview, error)
	ListArticleCreatedBy(ctx context.Context, userID string) ([]*ArticleOverview, error)
	TrendingArticle(ctx context.Context, p core.Pagination) ([]*ArticleOverview, error)
	CreateArticle(ctx context.Context, a *core.Article) (string, error)
	UpdateArticle(ctx context.Context, a *core.Article) error
	DeleteArticle(ctx context.Context, id string) error
	GetArticleDetails(ctx context.Context, id string) (*ArticleDetails, error)
	GetArticle(ctx context.Context, id string) (*core.Article, error)

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

//NewGW return a default implementation of GW service
func NewGW() (GW, error) {
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
	return articlesToArticleOverview(articles), nil
}

func (gw *gwService) ListArticleCreatedBy(ctx context.Context, userID string) ([]*ArticleOverview, error) {
	articles, err := gw.articleService.ListCreatedBy(ctx, userID)
	if err != nil {
		return []*ArticleOverview{}, err
	}
	return articlesToArticleOverview(articles), nil
}

func (gw *gwService) TrendingArticle(ctx context.Context, p core.Pagination) ([]*ArticleOverview, error) {
	articles, err := gw.ListArticle(context.Background(), core.Pagination{})
	if err != nil {
		return articles, err
	}
	sort.Slice(articles, func(i, j int) bool {
		if less := articles[j].VoteCount < articles[i].VoteCount; less {
			return true
		}
		if less := articles[i].VoteCount < articles[j].VoteCount; less {
			return false
		}
		return articles[i].LastUpdate.After(articles[j].LastUpdate)
	})
	return articles, nil
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

func (gw *gwService) GetArticleDetails(ctx context.Context, id string) (*ArticleDetails, error) {
	a, err := gw.articleService.Get(ctx, id)
	if err != nil {
		return &ArticleDetails{}, err
	}
	return &ArticleDetails{
		ArticleOverview: &ArticleOverview{
			Article: a,
		},
	}, nil
}

func (gw *gwService) GetArticle(ctx context.Context, id string) (*core.Article, error) {
	return gw.articleService.Get(ctx, id)
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
	id, err := gw.voteService.Create(ctx, v)
	if err == nil {
		gw.articleService.UpdateStatistic(ctx, v.ArticleID, &core.ArticleStatistic{VoteCount: 1})
	}
	return id, nil
}

func articlesToArticleOverview(articles []*core.Article) []*ArticleOverview {
	articlesOverview := make([]*ArticleOverview, 0)
	ids := make([]string, 0)
	for _, a := range articles {
		a := a
		articlesOverview = append(articlesOverview, &ArticleOverview{
			Article: a,
		})
		ids = append(ids, a.ID)
	}
	return articlesOverview
}

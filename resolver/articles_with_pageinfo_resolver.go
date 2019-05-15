package resolver

import (
	"context"

	"github.com/hatena/go-Intern-Diary/model"
)

type articlesWithPageInfoResolver struct {
	awp *model.ArticlesWithPageInfo
}

func (a *articlesWithPageInfoResolver) Articles(ctx context.Context) ([]*articleResolver, error) {
	articleResolvers := make([]*articleResolver, len(a.awp.Articles))
	for i, article := range a.awp.Articles {
		articleResolvers[i] = &articleResolver{article: article}
	}
	return articleResolvers, nil
}

func (a *articlesWithPageInfoResolver) PageInfo(ctx context.Context) (*pageInfoResolver, error) {
	return &pageInfoResolver{pageInfo: a.awp.PageInfo}, nil
}

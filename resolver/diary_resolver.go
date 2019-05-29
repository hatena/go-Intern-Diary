package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/loader"
	"github.com/hatena/go-Intern-Diary/model"
)

type diaryResolver struct {
	diary *model.Diary
	// app   service.DiaryApp
}

func (d *diaryResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(d.diary.ID))
}

func (d *diaryResolver) Name(ctx context.Context) string {
	return d.diary.Name
}

// func (d *diaryResolver) Articles(ctx context.Context) ([]*articleResolver, error) {
// 	page := uint64(1)
// 	limit := uint64(100)
// 	articles, err := d.app.ListArticlesByDiaryID(d.diary.ID, page, limit)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ars := make([]*articleResolver, len(articles))
// 	for i, article := range articles {
// 		ars[i] = &articleResolver{article}
// 	}
// 	return ars, nil
// }

func (d *diaryResolver) Articles(ctx context.Context) ([]*articleResolver, error) {
	articles, err := loader.LoadArticlesByDiaryID(ctx, d.diary.ID)
	if err != nil {
		return nil, err
	}
	articleResolvers := make([]*articleResolver, len(articles))
	for i, article := range articles {
		articleResolvers[i] = &articleResolver{article: article}
	}
	return articleResolvers, nil
}

func (d *diaryResolver) Tags(ctx context.Context) ([]*tagResolver, error) {
	tags, err := loader.LoadTagsByDiaryID(ctx, d.diary.ID)
	if err != nil {
		return nil, err
	}
	tagResolvers := make([]*tagResolver, len(tags))
	for i, tag := range tags {
		tagResolvers[i] = &tagResolver{tag: tag}
	}
	return tagResolvers, nil
}

func (d *diaryResolver) User(ctx context.Context) (*userResolver, error) {
	user, err := loader.LoadUserByDiaryID(ctx, d.diary.ID)
	if err != nil {
		return nil, err
	}
	return &userResolver{user: user}, nil
}

func (d *diaryResolver) CanEdit(ctx context.Context) bool {
	return d.diary.CanEdit
}

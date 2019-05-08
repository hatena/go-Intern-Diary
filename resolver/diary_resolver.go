package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

type diaryResolver struct {
	diary *model.Diary
	app   service.DiaryApp
}

func (d *diaryResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(d.diary.ID))
}

func (d *diaryResolver) Name(ctx context.Context) string {
	return d.diary.Name
}

func (d *diaryResolver) Articles(ctx context.Context) ([]*articleResolver, error) {
	page := uint64(1)
	limit := uint64(100)
	articles, err := d.app.ListArticlesByDiaryID(d.diary.ID, page, limit)
	if err != nil {
		return nil, err
	}
	ars := make([]*articleResolver, len(articles))
	for i, article := range articles {
		ars[i] = &articleResolver{article}
	}
	return ars, nil
}

package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/loader"
	"github.com/hatena/go-Intern-Diary/model"
)

type tagResolver struct {
	tag *model.Tag
}

func (t *tagResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(t.tag.ID))
}

func (t *tagResolver) TagName(ctx context.Context) string {
	return t.tag.TagName
}

func (t *tagResolver) CategoryID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(t.tag.CategoryID))
}

func (t *tagResolver) Diaries(ctx context.Context) ([]*diaryResolver, error) {
	diaries, err := loader.LoadDiariesByTagID(ctx, t.tag.ID)
	if err != nil {
		return nil, err
	}
	diairyResolvers := make([]*diaryResolver, len(diaries))
	for i, diary := range diaries {
		diairyResolvers[i] = &diaryResolver{diary: diary}
	}
	return diairyResolvers, nil
}

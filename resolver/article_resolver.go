package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/model"
)

type articleResolver struct {
	article *model.Article
}

func (a *articleResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(a.article.ID))
}

func (a *articleResolver) Title(ctx context.Context) string {
	return a.article.Title
}

func (a *articleResolver) Content(ctx context.Context) string {
	return a.article.Content
}

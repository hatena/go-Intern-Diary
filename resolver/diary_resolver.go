package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/model"
)

type diaryResolver struct {
	diary *model.Diary
}

func (d *diaryResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(d.diary.ID))
}

func (d *diaryResolver) Name(ctx context.Context) string {
	return d.diary.Name
}

// func (d *diaryResolver) Articles(ctx context.Context) []*articleResolver {

// }

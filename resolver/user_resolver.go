package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

type userResolver struct {
	user *model.User
	app  service.DiaryApp
}

func (u *userResolver) ID(ctx context.Context) graphql.ID {
	return graphql.ID(fmt.Sprint(u.user.ID))
}

func (u *userResolver) Name(ctx context.Context) string {
	return u.user.Name
}

func (u *userResolver) Diaries(ctx context.Context) ([]*diaryResolver, error) {

}

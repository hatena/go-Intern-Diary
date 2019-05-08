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
	page := uint64(1)
	limit := uint64(100) // todo
	diaries, err := u.app.ListDiariesByUserID(u.user.ID, page, limit)
	if err != nil {
		return nil, err
	}
	diaryResolvers := make([]*diaryResolver, len(diaries))
	for i, diary := range diaries {
		diaryResolvers[i] = &diaryResolver{diary, u.app}
	}
	return diaryResolvers, nil
}

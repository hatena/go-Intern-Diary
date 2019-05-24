package loader

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/dataloader"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

const userLoaderKey = "userLoader"

type userIDKey struct {
	id uint64
}

func (key userIDKey) String() string {
	return fmt.Sprint(key.id)
}

func (key userIDKey) Raw() interface{} {
	return key.id
}

func LoadUser(ctx context.Context, id uint64) (*model.User, error) {
	ldr, err := getLoader(ctx, userLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, userIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.(*model.User), nil
}

func LoadUserByDiaryID(ctx context.Context, id uint64) (*model.User, error) {
	ldr, err := getLoader(ctx, userLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, diaryIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.(*model.User), nil
}

func newUserLoader(app service.DiaryApp) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		userIDs := make([]uint64, len(keys))
		diaryIDs := make([]uint64, len(keys))
		for _, key := range keys {
			switch key := key.(type) {
			case userIDKey:
				userIDs = append(userIDs, key.id)
			case diaryIDKey:
				diaryIDs = append(diaryIDs, key.id)
			}
		}
		usersByIDs, _ := app.ListUsersByIDs(userIDs)
		usersByDiaryIDs, _ := app.ListUsersByDiaryIDs(diaryIDs)
		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}
			switch key := key.(type) {
			case userIDKey:
				for _, user := range usersByIDs {
					if key.id == user.ID {
						results[i].Data = user
						continue
					}
				}
				if results[i].Data == nil {
					results[i].Error = errors.New("user not found")
				}
			case diaryIDKey:
				results[i].Data = usersByDiaryIDs[key.id]
			}
		}
		return results
	}
}

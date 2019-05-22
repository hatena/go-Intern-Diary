package loader

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/dataloader"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

const diaryLoaderKey = "diaryLoader"

type diaryIDKey struct {
	id uint64
}

func (key diaryIDKey) String() string {
	return fmt.Sprint(key.id)
}

func (key diaryIDKey) Raw() interface{} {
	return key.id
}

func LoadDiary(ctx context.Context, id uint64) (*model.Diary, error) {
	ldr, err := getLoader(ctx, diaryLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, diaryIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.(*model.Diary), nil
}

func LoadDiariesByUserID(ctx context.Context, id uint64) ([]*model.Diary, error) {
	ldr, err := getLoader(ctx, diaryLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, userIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.([]*model.Diary), nil
}

func LoadDiariesByTagID(ctx context.Context, id uint64) ([]*model.Diary, error) {
	ldr, err := getLoader(ctx, diaryLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, tagIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.([]*model.Diary), nil
}

func newDiaryLoader(app service.DiaryApp) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		diaryIDs := make([]uint64, 0, len(keys))
		userIDs := make([]uint64, 0, len(keys))
		tagIDs := make([]uint64, 0, len(keys))
		for _, key := range keys {
			switch key := key.(type) {
			case diaryIDKey:
				diaryIDs = append(diaryIDs, key.id)
			case userIDKey:
				userIDs = append(userIDs, key.id)
			case tagIDKey:
				tagIDs = append(tagIDs, key.id)
			}
		}
		diaries, _ := app.ListDiariesByIDs(diaryIDs)
		diariesByUserIDs, _ := app.ListDiariesByUserIDs(userIDs)
		diariesByTagIDs, _ := app.ListDiariesByTagIDs(tagIDs)
		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}
			switch key := key.(type) {
			case diaryIDKey:
				for _, diary := range diaries {
					if key.id == diary.ID {
						results[i].Data = diary
						continue
					}
				}
				if results[i].Data == nil {
					results[i].Error = errors.New("diary not found")
				}
			case userIDKey:
				results[i].Data = diariesByUserIDs[key.id]
			case tagIDKey:
				results[i].Data = diariesByTagIDs[key.id]
			}
		}
		return results
	}
}

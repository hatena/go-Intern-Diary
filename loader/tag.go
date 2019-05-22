package loader

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

const tagLoaderKey = "tagLoader"

type tagIDKey struct {
	id uint64
}

func (key tagIDKey) String() string {
	return fmt.Sprint(key.id)
}

func (key tagIDKey) Raw() interface{} {
	return key.id
}

func LoadTag(ctx context.Context, id uint64) (*model.Tag, error) {
	ldr, err := getLoader(ctx, tagLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, tagIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.(*model.Tag), nil
}

func LoadTagsByDiaryID(ctx context.Context, id uint64) ([]*model.Tag, error) {
	ldr, err := getLoader(ctx, tagLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, diaryIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.([]*model.Tag), nil
}

func newTagLoader(app service.DiaryApp) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		tagIDs := make([]uint64, 0, len(keys))
		diaryIDs := make([]uint64, 0, len(keys))
		for _, key := range keys {
			switch key := key.(type) {
			case tagIDKey:
				tagIDs = append(tagIDs, key.id)
			case diaryIDKey:
				diaryIDs = append(diaryIDs, key.id)
			}
		}
		tagsByIDs, _ := app.ListTagsByIDs(tagIDs)
		tagsByDiaryIDs, _ := app.ListTagsByDiaryIDs(diaryIDs)
		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}
			switch key := key.(type) {
			case tagIDKey:
				for _, tag := range tagsByIDs {
					if key.id == tag.ID {
						results[i].Data = tag
						continue
					}
				}
				if results[i].Data == nil {
					results[i].Error = errors.New("tag not found")
				}
			case diaryIDKey:
				results[i].Data = tagsByDiaryIDs[key.id]
			}
		}
		return results
	}
}

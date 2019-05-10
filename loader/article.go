package loader

import (
	"context"
	"errors"
	"fmt"

	"github.com/graph-gophers/dataloader"
	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/service"
)

const articleLoaderKey = "articleLoader"

type articleIDKey struct {
	id uint64
}

func (key articleIDKey) String() string {
	return fmt.Sprint(key.id)
}

func (key articleIDKey) Raw() interface{} {
	return key.id
}

func LoadArticle(ctx context.Context, id uint64) (*model.Article, error) {
	ldr, err := getLoader(ctx, articleLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, articleIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.(*model.Article), nil
}

func LoadArticlesByDiaryID(ctx context.Context, id uint64) ([]*model.Article, error) {
	ldr, err := getLoader(ctx, articleLoaderKey)
	if err != nil {
		return nil, err
	}
	data, err := ldr.Load(ctx, diaryIDKey{id: id})()
	if err != nil {
		return nil, err
	}
	return data.([]*model.Article), nil
}

func newArticleLoader(app service.DiaryApp) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		articleIDs := make([]uint64, 0, len(keys))
		diaryIDs := make([]uint64, 0, len(keys))
		for _, key := range keys {
			switch key := key.(type) {
			case articleIDKey:
				articleIDs = append(articleIDs, key.id)
			case diaryIDKey:
				diaryIDs = append(diaryIDs, key.id)
			}
		}
		articlesByIDs, _ := app.ListArticlesByIDs(articleIDs)
		articlesByDiaryIDs, _ := app.ListArticlesByDiaryIDs(diaryIDs)
		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}
			switch key := key.(type) {
			case articleIDKey:
				for _, article := range articlesByIDs {
					if key.id == article.ID {
						results[i].Data = article
						continue
					}
				}
				if results[i].Data == nil {
					results[i].Error = errors.New("article not found")
				}
			case diaryIDKey:
				results[i].Data = articlesByDiaryIDs[key.id]
			}
		}
		return results
	}
}

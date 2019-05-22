package service

import "github.com/hatena/go-Intern-Diary/model"

func (app *diaryApp) ListTagsByIDs(tagIDs []uint64) ([]*model.Tag, error) {
	return app.repo.ListTagsByIDs(tagIDs)
}

func (app *diaryApp) ListTagsByDiaryIDs(diaryIDs []uint64) (map[uint64][]*model.Tag, error) {
	return app.repo.ListTagsByDiaryIDs(diaryIDs)
}

package service

import (
	"errors"

	"github.com/hatena/go-Intern-Diary/model"
)

func (app *diaryApp) CreateNewDiary(userID uint64, name string) (*model.Diary, error) {
	return app.repo.CreateNewDiary(userID, name)
}

func (app *diaryApp) ListDiariesByUserID(userID, page, limit uint64) ([]*model.Diary, error) {
	if page < 1 || limit < 1 {
		return nil, errors.New("page and limit should be positive")
	}
	return app.repo.ListDiariesByUserID(userID, limit, (page-1)*limit)
}

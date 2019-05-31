package service

import (
	"errors"

	"github.com/hatena/go-Intern-Diary/model"
)

func (app *diaryApp) CreateNewDiary(userID uint64, name string, tagWithCategories []*model.TagWithCategory) (*model.Diary, error) {
	return app.repo.CreateNewDiary(userID, name, tagWithCategories)
}

func (app *diaryApp) ListDiariesByUserID(userID, page, limit uint64) ([]*model.Diary, error) {
	if page < 1 || limit < 1 {
		return nil, errors.New("page and limit should be positive")
	}
	return app.repo.ListDiariesByUserID(userID, limit, (page-1)*limit)
}

func (app *diaryApp) DeleteDiary(userID, diaryID uint64) error {
	return app.repo.DeleteDiary(userID, diaryID)
}

func (app *diaryApp) FindDiaryByID(diaryID uint64) (*model.Diary, error) {
	return app.repo.FindDiaryByID(diaryID)
}

func (app *diaryApp) ListDiariesByIDs(diaryIDs []uint64) ([]*model.Diary, error) {
	return app.repo.ListDiariesByIDs(diaryIDs)
}

func (app *diaryApp) ListDiariesByUserIDs(userIDs []uint64) (map[uint64][]*model.Diary, error) {
	return app.repo.ListDiariesByUserIDs(userIDs)
}

func (app *diaryApp) ListDiariesByTagIDs(tagIDs []uint64) (map[uint64][]*model.Diary, error) {
	return app.repo.ListDiariesByTagIDs(tagIDs)
}

func (app *diaryApp) ListRecommendedDiaries(diaryID uint64) ([]*model.Diary, error) {
	return app.repo.ListRecommendedDiaries(diaryID)
}

func (app *diaryApp) ListCategories() []*model.Category {
	return app.repo.ListCategories()
}

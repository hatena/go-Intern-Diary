package service

import (
	"errors"

	"github.com/hatena/go-Intern-Diary/model"
)

func (app *diaryApp) ListArticlesByDiaryID(diaryID, page, limit uint64) ([]*model.Article, error) {
	if page < 1 || limit < 1 {
		return nil, errors.New("page and limit should be positive")
	}
	return app.repo.ListArticlesByDiaryID(diaryID, limit, (page-1)*limit)
}

func (app *diaryApp) CreateNewArticle(diaryID uint64, title string, content string) (*model.Article, error) {
	return app.repo.CreateNewArticle(diaryID, title, content)
}

func (app *diaryApp) FindArticleByID(articleID, diaryID uint64) (*model.Article, error) {
	return app.repo.FindArticleByID(articleID, diaryID)
}
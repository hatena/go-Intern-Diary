package service

import (
	"errors"

	"github.com/hatena/go-Intern-Diary/model"
)

func (app *diaryApp) ListArticlesByDiaryID(diaryID uint64, page, limit int) ([]*model.Article, *model.PageInfo, error) {
	if page < 1 || limit < 1 {
		return nil, nil, errors.New("page and limit should be positive")
	}
	articles, pageInfo, err := app.repo.ListArticlesByDiaryID(diaryID, page, limit) // スタートが1
	if err != nil {
		return nil, nil, err
	}
	if pageInfo.CurrentPage > pageInfo.TotalPage {
		return nil, nil, errors.New("Invalid Page Paramter")
	}
	return articles, pageInfo, nil
}

func (app *diaryApp) CreateNewArticle(diaryID uint64, title string, content string) (*model.Article, error) {
	return app.repo.CreateNewArticle(diaryID, title, content)
}

func (app *diaryApp) FindArticleByID(articleID, diaryID uint64) (*model.Article, error) {
	return app.repo.FindArticleByID(articleID, diaryID)
}

func (app *diaryApp) DeleteArticle(articleID uint64) error {
	return app.repo.DeleteArticle(articleID)
}

func (app *diaryApp) ListArticlesByIDs(articleIDs []uint64) ([]*model.Article, error) {
	return app.repo.ListArticlesByIDs(articleIDs)
}

func (app *diaryApp) ListArticlesByDiaryIDs(diaryIDs []uint64) (map[uint64][]*model.Article, error) {
	return app.repo.ListArticlesByDiaryIDs(diaryIDs)
}

func (app *diaryApp) UpdateArticle(articleID uint64, title, content string) (*model.Article, error) {
	return app.repo.UpdateArticle(articleID, title, content)
}

package service

import (
	"math/rand"
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/hatena/go-Intern-Diary/repository"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DiaryApp interface {
	CreateNewUser(name string, passwordHash string) error
	FindUserByName(name string) (*model.User, error)
	FindUserByID(id uint64) (*model.User, error)
	ListUsersByIDs(userIDs []uint64) ([]*model.User, error)
	LoginUser(name string, password string) (bool, error)
	CreateNewToken(userID uint64, expiresAt time.Time) (string, error)
	FindUserByToken(token string) (*model.User, error)

	CreateNewDiary(userID uint64, name string) (*model.Diary, error)
	ListDiariesByUserID(userID, page, limit uint64) ([]*model.Diary, error)
	DeleteDiary(userID, diaryID uint64) error
	ListArticlesByDiaryID(diaryID, page, limit uint64) ([]*model.Article, error)
	FindDiaryByID(diaryID, userID uint64) (*model.Diary, error)
	CreateNewArticle(diaryID uint64, title string, content string) (*model.Article, error)
	FindArticleByID(articleID, diaryID uint64) (*model.Article, error)
	DeleteArticle(articleID, diaryID uint64) error

	Close() error
}

func NewApp(repo repository.Repository) DiaryApp {
	return &diaryApp{repo: repo}
}

type diaryApp struct {
	repo repository.Repository
}

func (app *diaryApp) Close() error {
	return app.repo.Close()
}

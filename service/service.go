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
	ListUsersByDiaryIDs(diaryIDs []uint64) (map[uint64]*model.User, error)
	LoginUser(name string, password string) (bool, error)
	CreateNewToken(userID uint64, expiresAt time.Time) (string, error)
	FindUserByToken(token string) (*model.User, error)

	CreateNewDiary(userID uint64, name string, tagWithCategories []*model.TagWithCategory) (*model.Diary, error)
	ListDiariesByUserID(userID, page, limit uint64) ([]*model.Diary, error)
	DeleteDiary(userID, diaryID uint64) error
	ListArticlesByDiaryID(diaryID uint64, page, limit int) ([]*model.Article, *model.PageInfo, error)
	FindDiaryByID(diaryID uint64) (*model.Diary, error)
	CreateNewArticle(diaryID, userID uint64, title string, content string) (*model.Article, error)
	FindArticleByID(articleID, diaryID uint64) (*model.Article, error)
	UpdateArticle(articleID, userID uint64, title, content string) (*model.Article, error)
	DeleteArticle(articleID, userID uint64) error

	ListDiariesByIDs(diaryIDs []uint64) ([]*model.Diary, error)
	ListDiariesByUserIDs(userIDs []uint64) (map[uint64][]*model.Diary, error)

	ListArticlesByIDs(articleIDs []uint64) ([]*model.Article, error)
	ListArticlesByDiaryIDs(diaryIDs []uint64) (map[uint64][]*model.Article, error)

	ListDiariesByTagIDs(tagIDs []uint64) (map[uint64][]*model.Diary, error)
	ListTagsByIDs(tagIDs []uint64) ([]*model.Tag, error)
	ListTagsByDiaryIDs(diaryIDs []uint64) (map[uint64][]*model.Tag, error)

	ListRecommendedDiaries(diaryID uint64) ([]*model.Diary, error)
	ListCategories() []*model.Category

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

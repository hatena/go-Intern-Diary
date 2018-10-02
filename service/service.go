package service

import (
	"math/rand"
	"time"

	"github.com/hatena/go-Intern-Diary/repository"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DiaryApp interface {
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

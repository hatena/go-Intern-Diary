package repository

import (
	"time"

	"github.com/hatena/go-Intern-Diary/model"
)

var diaryNotFoundError = model.NotFoundError("diary")

func (r *repository) CreateNewDiary(userID uint64, name string) (*model.Diary, error) {
	id, err := r.generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO diary
			(id, name, user_id, created_at, updated_at)
			VALUES(?, ?, ?, ?, ?)
			`, id, name, userID, now, now,
	)
	if err != nil {
		return nil, err
	}
	return &model.Diary{ID: id, Name: name, UserID: userID, UpdatedAt: now}, nil
}

func (r *repository) ListDiariesByUserID(userID uint64, limit, offset uint64) ([]*model.Diary, error) {
	diaries := make([]*model.Diary, 0, limit)
	err := r.db.Select(&diaries,
		`SELECT id, name, user_id, updated_at FROM diary
			WHERE user_id = ?
			ORDER BY updated_at DESC LIMIT ? OFFSET ?`,
		userID, limit, offset,
	)
	return diaries, err
}

package repository

import (
	"database/sql"
	"time"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
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
			ORDER BY created_at DESC LIMIT ? OFFSET ?`,
		userID, limit, offset,
	)
	return diaries, err
}

func (r *repository) DeleteDiary(userID, diaryID uint64) (err error) {
	_, err = r.db.Exec(
		`DELETE FROM diary
			WHERE user_id = ? AND id = ?`,
		userID, diaryID,
	)
	if err != nil {
		return
	}
	_, err = r.db.Exec(
		`DELETE FROM article
			WHERE diary_id = ?`,
		diaryID,
	)
	return
}

func (r *repository) FindDiaryByID(diaryID, userID uint64) (*model.Diary, error) {
	var diary model.Diary
	err := r.db.Get(
		&diary,
		`SELECT id, name FROM diary
			WHERE id = ? AND user_id = ? LIMIT 1`,
		diaryID, userID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, diaryNotFoundError
		}
		return nil, err
	}
	return &diary, nil
}

func (r *repository) ListDiariesByIDs(diaryIDs []uint64) ([]*model.Diary, error) {
	if len(diaryIDs) == 0 {
		return nil, nil
	}
	diaries := make([]*model.Diary, 0, len(diaryIDs))
	query, args, err := sqlx.In(
		`SELECT id, name, user_id, updated_at FROM diary
			WHERE id IN (?)
			ORDER BY created_at DESC`, diaryIDs,
	)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&diaries, query, args...)
	return diaries, err
}

func (r *repository) ListDiariesByUserIDs(userIDs []uint64) (map[uint64][]*model.Diary, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(
		`SELECT id, name, user_id, updated_at FROM diary
		WHERE user_id IN (?)
		ORDER BY created_at DESC`,
		userIDs,
	)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	diaries := make(map[uint64][]*model.Diary)
	for rows.Next() {
		var diary model.Diary
		rows.Scan(&diary.ID, &diary.Name, &diary.UserID, &diary.UpdatedAt)
		diaries[diary.UserID] = append(diaries[diary.UserID], &diary)
	}
	return diaries, nil
}

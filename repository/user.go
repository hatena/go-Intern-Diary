package repository

import (
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hatena/go-Intern-Diary/model"
	"github.com/jmoiron/sqlx"
)

var userNotFoundError = model.NotFoundError("user")

func (r *repository) CreateNewUser(name string, passwordHash string) error {
	id, err := r.generateID()
	if err != nil {
		return err
	}

	now := time.Now()
	_, err = r.db.Exec(
		`INSERT INTO user
			(id, name, password_hash, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)
		`,
		id, name, passwordHash, now, now,
	)
	return err
}

func (r *repository) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := r.db.Get(
		&user,
		`SELECT id, name FROM user 
				WHERE name = ? LIMIT 1
			`, name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userNotFoundError
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindUserByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.Get(
		&user,
		`SELECT id, name FROM user
			WHERE id = ? LIMIT 1
		`, id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userNotFoundError
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) ListUsersByIDs(userIDs []uint64) ([]*model.User, error) {
	if len(userIDs) == 0 {
		return nil, nil
	}
	users := make([]*model.User, 0, len(userIDs))
	query, args, err := sqlx.In(
		`SELECT id, name FROM user
			WHERE id IN (?)`, userIDs,
	)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&users, query, args...)
	return users, err
}

func (r *repository) FindPasswordHashByName(name string) (string, error) {
	var hash string
	err := r.db.Get(
		&hash,
		`SELECT password_hash FROM user
			WHERE name = ? LIMIT 1`, name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
	}
	return hash, nil
}

func (r *repository) CreateNewToken(userID uint64, token string, expiresAt time.Time) error {
	now := time.Now()
	_, err := r.db.Exec(
		`INSERT INTO user_session
			(user_id, token, expires_at, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?)`,
		userID, token, expiresAt, now, now,
	)
	return err
}

func (r *repository) FindUserByToken(token string) (*model.User, error) {
	var user model.User
	err := r.db.Get(
		&user,
		`SELECT id, name FROM user 
			JOIN user_session ON user.id = user_session.user_id
			WHERE user_session.token = ? && user_session.expires_at > ?
			LIMIT 1`, token, time.Now(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, userNotFoundError
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) ListUsersByDiaryIDs(diaryIDs []uint64) (map[uint64]*model.User, error) {
	if len(diaryIDs) == 0 {
		return nil, nil
	}
	type UserWithDiaryID struct {
		ID      uint64 `db:"id"`
		Name    string `db:"name"`
		DiaryID uint64 `db:"diary_id"`
	}
	userWithDiaryIDs := make([]*UserWithDiaryID, 0, len(diaryIDs))
	query, args, err := sqlx.In(
		`SELECT user.id as id, user.name as name, diary.id as diary_id FROM user
			JOIN diary ON user.id = diary.user_id
			WHERE diary.id IN (?)
			`, diaryIDs,
	)
	if err != nil {
		return nil, err
	}
	err = r.db.Select(&userWithDiaryIDs, query, args...)
	if err != nil {
		return nil, err
	}
	users := make(map[uint64]*model.User, len(diaryIDs))
	for _, u := range userWithDiaryIDs {
		users[u.DiaryID] = &model.User{
			ID:   u.ID,
			Name: u.Name,
		}
	}
	return users, nil
}

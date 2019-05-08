package model

import "time"

type Article struct {
	ID        uint64    `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	DiaryID   uint64    `db:"diary_id"`
	UpdatedAt time.Time `db:"updated_at"`
}

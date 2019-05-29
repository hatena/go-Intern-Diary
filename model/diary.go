package model

import "time"

type Diary struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	UserID    uint64    `db:"user_id"`
	UpdatedAt time.Time `db:"updated_at"`
	CanEdit   bool
}

type RecommendedDiary struct {
	DiaryID   uint64
	DiaryName string
	UpdatedAt time.Time
	UserID    uint64
	UserName  string
}

package model

import "time"

type Diary struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	UserID    uint64    `db:"user_id"`
	UpdatedAt time.Time `db:"updated_at"`
}

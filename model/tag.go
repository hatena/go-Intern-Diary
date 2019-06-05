package model

import "database/sql"

type Tag struct {
	ID         uint64 `db:"id"`
	TagName    string `db:"tag_name"`
	CategoryID int    `db:"category_id"`
}

type TagRecord struct {
	ID         uint64        `db:"id"`
	TagName    string        `db:"tag_name"`
	CategoryID sql.NullInt64 `db:"category_id"`
}

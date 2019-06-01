package model

type Tag struct {
	ID         uint64 `db:"id"`
	TagName    string `db:"tag_name"`
	CategoryID int    `db:"category_id"`
}

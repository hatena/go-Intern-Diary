package model

type Category struct {
	ID           int    `db:"id"`
	CategoryName string `db:"category_name"`
}

package model

type Tag struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type TagBook struct {
	ID     int64 `db:"id" json:"id"`
	TagID  int64 `db:"tag_id" json:"tag_id"`
	BookID int64 `db:"book_id" json:"book_id"`
}

type TagBookForIndex struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	BookID int64  `db:"book_id" json:"book_id"`
}

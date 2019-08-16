package model

type Book struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ISBN     int64  `db:"isbn" json:"isbn"`
	ImageURL string `db:"image_url" json:"image_url"`
}

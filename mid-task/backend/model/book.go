package model

type Book struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ISBN     string `db:"isbn" json:"isbn"`
	ImageURL string `db:"image_url" json:"image_url"`
	ItemURL  string `db:"item_url" json:"item_url"`
}

type BookWithTags struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	ISBN     string `db:"isbn" json:"isbn"`
	ImageURL string `db:"image_url" json:"image_url"`
	ItemURL  string `db:"item_url" json:"item_url"`
	Tags     []Tag  `json:"tags"`
}
type SearchedBooks struct {
	Items []struct {
		Item struct {
			Title         string `json:"title"`
			Author        string `json:"author"`
			Isbn          string `json:"isbn"`
			ItemURL       string `json:"itemUrl"`
			LargeImageURL string `json:"largeImageUrl"`
		} `json:"Item"`
	} `json:"Items"`
}

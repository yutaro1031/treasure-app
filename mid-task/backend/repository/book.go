package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
)

func AllBook(db *sqlx.DB) ([]model.BookWithTags, error) {
	b := make([]model.BookWithTags, 0)
	if err := db.Select(&b, `SELECT id, name, isbn, image_url, item_url FROM book`); err != nil {
		return nil, err
	}
	return b, nil
}

func FindBook(db *sqlx.DB, id int64) (*model.Book, error) {
	b := model.Book{}
	if err := db.Get(&b, `
SELECT id, name, isbn, image_url, item_url FROM book WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &b, nil
}

func CreateBook(db *sqlx.Tx, b *model.Book) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO book (name, isbn, image_url, item_url) VALUES (?, ?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(b.Name, b.ISBN, b.ImageURL, b.ItemURL)
}

func DestroyBook(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM book WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}

package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
)

func AllBook(db *sqlx.DB) ([]model.Book, error) {
	a := make([]model.Book, 0)
	if err := db.Select(&a, `SELECT id, name, isbn, image_url, item_url FROM book`); err != nil {
		return nil, err
	}
	return a, nil
}

func FindBook(db *sqlx.DB, id int64) (*model.Book, error) {
	a := model.Book{}
	if err := db.Get(&a, `
SELECT id, name, isbn, image_url, item_url FROM book WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateBook(db *sqlx.Tx, a *model.Book) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO book (name, isbn, image_url, item_url) VALUES (?, ?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Name, a.ISBN, a.ImageURL, a.ItemURL)
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

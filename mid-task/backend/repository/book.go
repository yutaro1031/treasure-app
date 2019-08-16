package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllBook(db *sqlx.DB) ([]model.Book, error) {
	a := make([]model.Book, 0)
	if err := db.Select(&a, `SELECT id, name, isbn, image_url FROM book`); err != nil {
		return nil, err
	}
	return a, nil
}

// func FindBook(db *sqlx.DB, id int64) (*model.Book, error) {
// 	a := model.Book{}
// 	if err := db.Get(&a, `
// SELECT id, title, body FROM book WHERE id = ?
// `, id); err != nil {
// 		return nil, err
// 	}
// 	return &a, nil
// }

func CreateBook(db *sqlx.Tx, a *model.Book) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO book (name, isbn, image_url) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Name, a.ISBN, a.ImageURL)
}

// func UpdateBook(db *sqlx.Tx, id int64, a *model.Book) (sql.Result, error) {
// 	stmt, err := db.Prepare(`
// UPDATE book SET title = ?, body = ? WHERE id = ?
// `)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	return stmt.Exec(a.Title, a.Body, id)
// }

// func DestroyBook(db *sqlx.Tx, id int64) (sql.Result, error) {
// 	stmt, err := db.Prepare(`
// DELETE FROM book WHERE id = ?
// `)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()
// 	return stmt.Exec(id)
// }

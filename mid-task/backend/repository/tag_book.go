package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
)

func AllTagBook(db *sqlx.DB) ([]model.TagBookForIndex, error) {
	b := make([]model.TagBookForIndex, 0)
	if err := db.Select(&b, `
SELECT tag_book.id AS 'id', tag.name AS 'name', book.id AS 'book_id'
FROM tag_book INNER JOIN tag ON tag_book.tag_id = tag.id
INNER JOIN book ON tag_book.book_id = book.id
`); err != nil {
		return nil, err
	}
	return b, nil
}

func FindTagBook(db *sqlx.DB, id int64) (*model.TagBook, error) {
	t := model.TagBook{}
	if err := db.Get(&t, `
SELECT id, tag_id, book_id FROM tag_book WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &t, nil
}

func CreateTagBook(db *sqlx.Tx, t *model.TagBook) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO tag_book (tag_id, book_id) VALUES (?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.TagID, t.BookID)
}

func DestroyTagBook(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM tag_book WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}

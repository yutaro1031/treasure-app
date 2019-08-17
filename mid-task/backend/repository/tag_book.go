package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
)

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

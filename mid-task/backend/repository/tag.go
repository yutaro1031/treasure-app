package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
)

func AllTag(db *sqlx.DB) ([]model.Tag, error) {
	t := make([]model.Tag, 0)
	if err := db.Select(&t, `SELECT id, name FROM tag`); err != nil {
		return nil, err
	}
	return t, nil
}

func FindTag(db *sqlx.DB, id int64) (*model.Tag, error) {
	t := model.Tag{}
	if err := db.Get(&t, `
SELECT id, name FROM tag WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &t, nil
}

func CreateTag(db *sqlx.Tx, t *model.Tag) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO tag (name) VALUES (?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.Name)
}

func UpdateTag(db *sqlx.Tx, id int64, t *model.Tag) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE tag SET name = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(t.Name, id)
}

func DestroyTag(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM tag WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}

package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func CreateArticleComment(db *sqlx.Tx, ac *model.ArticleComment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_comment (user_id, article_id ,body) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(ac.UserID, ac.ArticleID, ac.Body)
}

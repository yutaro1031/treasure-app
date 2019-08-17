package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/yutaro1031/treasure-app/mid-task/backend/dbutil"

	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
	"github.com/yutaro1031/treasure-app/mid-task/backend/repository"
)

type TagBook struct {
	db *sqlx.DB
}

func NewTagBook(db *sqlx.DB) *TagBook {
	return &TagBook{db}
}

func (a *TagBook) Create(newTagBook *model.TagBook) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateTagBook(tx, newTagBook)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed tagBook insert transaction")
	}
	return createdId, nil
}

func (a *TagBook) Destroy(id int64) error {
	_, err := repository.FindTagBook(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find tagBook")
	}

	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyTagBook(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed tagBook delete transaction")
	}
	return nil
}

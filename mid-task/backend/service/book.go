package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/yutaro1031/treasure-app/mid-task/backend/dbutil"

	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
	"github.com/yutaro1031/treasure-app/mid-task/backend/repository"
)

type Book struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) *Book {
	return &Book{db}
}

// func (a *Book) Update(id int64, newBook *model.Book) error {
// 	_, err := repository.FindBook(a.db, id)
// 	if err != nil {
// 		return errors.Wrap(err, "failed find book")
// 	}

// 	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
// 		_, err := repository.UpdateBook(tx, id, newBook)
// 		if err != nil {
// 			return err
// 		}
// 		if err := tx.Commit(); err != nil {
// 			return err
// 		}
// 		return err
// 	}); err != nil {
// 		return errors.Wrap(err, "failed book update transaction")
// 	}
// 	return nil
// }

// func (a *Book) Destroy(id int64) error {
// 	_, err := repository.FindBook(a.db, id)
// 	if err != nil {
// 		return errors.Wrap(err, "failed find book")
// 	}

// 	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
// 		_, err := repository.DestroyBook(tx, id)
// 		if err != nil {
// 			return err
// 		}
// 		if err := tx.Commit(); err != nil {
// 			return err
// 		}
// 		return err
// 	}); err != nil {
// 		return errors.Wrap(err, "failed book delete transaction")
// 	}
// 	return nil
// }

func (a *Book) Create(newBook *model.Book) (int64, error) {
	var createdId int64
	if err := dbutil.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateBook(tx, newBook)
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
		return 0, errors.Wrap(err, "failed book insert transaction")
	}
	return createdId, nil
}

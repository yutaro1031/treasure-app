package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/yutaro1031/treasure-app/mid-task/backend/httputil"
	"github.com/yutaro1031/treasure-app/mid-task/backend/model"
	"github.com/yutaro1031/treasure-app/mid-task/backend/service"
)

type TagBook struct {
	db *sqlx.DB
}

func NewTagBook(db *sqlx.DB) *TagBook {
	return &TagBook{db: db}
}

func (a *TagBook) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newTagBook := &model.TagBook{}
	if err := json.NewDecoder(r.Body).Decode(&newTagBook); err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagBookService := service.NewTagBook(a.db)
	id, err := tagBookService.Create(newTagBook)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newTagBook.ID = id

	return http.StatusCreated, newTagBook, nil
}

func (a *TagBook) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagBookService := service.NewTagBook(a.db)
	err = tagBookService.Destroy(aid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

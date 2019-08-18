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

type Book struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) *Book {
	return &Book{db: db}
}

func (a *Book) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	bookService := service.NewBook(a.db)
	books, err := bookService.AllBookWithTags()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, books, nil
}

func (a *Book) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newBook := &model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		return http.StatusBadRequest, nil, err
	}

	bookService := service.NewBook(a.db)
	id, err := bookService.Create(newBook)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newBook.ID = id

	return http.StatusCreated, newBook, nil
}

func (a *Book) Search(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	param := r.URL.Query()
	keyword, ok := param["keyword"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	page, ok := param["page"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	searchedBooks, err := httputil.Search(keyword[0], page[0])
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, searchedBooks, nil
}

func (a *Book) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	bookService := service.NewBook(a.db)
	err = bookService.Destroy(aid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

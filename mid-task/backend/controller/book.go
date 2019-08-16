package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
)

type Book struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) *Book {
	return &Book{db: db}
}

func (a *Book) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	books, err := repository.AllBook(a.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, books, nil
}

// func (a *Book) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// 	vars := mux.Vars(r)
// 	id, ok := vars["id"]
// 	if !ok {
// 		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
// 	}

// 	aid, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}

// 	book, err := repository.FindBook(a.db, aid)
// 	if err != nil && err == sql.ErrNoRows {
// 		return http.StatusNotFound, nil, err
// 	} else if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}

// 	return http.StatusCreated, book, nil
// }

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

// func (a *Book) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// 	vars := mux.Vars(r)
// 	id, ok := vars["id"]
// 	if !ok {
// 		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
// 	}

// 	aid, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}

// 	reqBook := &model.Book{}
// 	if err := json.NewDecoder(r.Body).Decode(&reqBook); err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}

// 	bookService := service.NewBook(a.db)
// 	err = bookService.Update(aid, reqBook)
// 	if err != nil && errors.Cause(err) == sql.ErrNoRows {
// 		return http.StatusNotFound, nil, err
// 	} else if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}

// 	return http.StatusNoContent, nil, nil
// }

// func (a *Book) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
// 	vars := mux.Vars(r)
// 	id, ok := vars["id"]
// 	if !ok {
// 		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
// 	}

// 	aid, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return http.StatusBadRequest, nil, err
// 	}

// 	bookService := service.NewBook(a.db)
// 	err = bookService.Destroy(aid)
// 	if err != nil && errors.Cause(err) == sql.ErrNoRows {
// 		return http.StatusNotFound, nil, err
// 	} else if err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}

// 	return http.StatusNoContent, nil, nil
// }

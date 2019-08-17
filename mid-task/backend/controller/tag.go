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
	"github.com/yutaro1031/treasure-app/mid-task/backend/repository"
	"github.com/yutaro1031/treasure-app/mid-task/backend/service"
)

type Tag struct {
	db *sqlx.DB
}

func NewTag(db *sqlx.DB) *Tag {
	return &Tag{db: db}
}

func (a *Tag) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	tags, err := repository.AllTag(a.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, tags, nil
}

func (a *Tag) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagService := service.NewTag(a.db)
	id, err := tagService.Create(newTag)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newTag.ID = id

	return http.StatusCreated, newTag, nil
}

func (a *Tag) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqTag := &model.Tag{}
	if err := json.NewDecoder(r.Body).Decode(&reqTag); err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagService := service.NewTag(a.db)
	err = tagService.Update(aid, reqTag)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (a *Tag) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	tagService := service.NewTag(a.db)
	err = tagService.Destroy(aid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

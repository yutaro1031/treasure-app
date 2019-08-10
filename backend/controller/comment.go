package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/service"
)

type ArticleComment struct {
	dbx *sqlx.DB
}

func NewArticleComment(dbx *sqlx.DB) *ArticleComment {
	return &ArticleComment{dbx: dbx}
}

func (ac *ArticleComment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["article_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	createArticleComment := &model.CreateArticleComment{}
	if err := json.NewDecoder(r.Body).Decode(&createArticleComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	articleComment := &model.ArticleComment{
		ArticleID: aid,
		UserID:    user.ID,
		Body:      createArticleComment.Body,
	}

	articleCommentService := service.NewArticleCommentService(ac.dbx)
	createdID, err := articleCommentService.Create(articleComment)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	articleComment.ID = createdID
	return articleComment, http.StatusCreated, nil
}

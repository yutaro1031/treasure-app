package controller

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/service"
	"net/http"
)

type Comment struct {
	dbx *sqlx.DB
}

func (a *Comment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newArticleComment := &model.CreateArticleComment{}
	if err := json.NewDecoder(r.Body).Decode(&newArticleComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	user.
	articleCommentService := service.NewArticleCommentService(a.dbx)
	articleCommentService.Create()
}

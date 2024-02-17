package handlers

import (
	"github.com/makarellav/bloagg/internal/entity"
	"github.com/makarellav/bloagg/internal/http/response"
	"github.com/makarellav/bloagg/internal/services"
	"net/http"
	"strconv"
)

type PostHandlers struct {
	postSrv *services.PostService
}

func NewPostHandlers(postSrv *services.PostService) *PostHandlers {
	return &PostHandlers{postSrv: postSrv}
}

func (ph *PostHandlers) HandleGetByUserID(w http.ResponseWriter, r *http.Request, user *entity.User) {
	limitQuery := r.URL.Query().Get("limit")

	limit, err := strconv.Atoi(limitQuery)

	if err != nil || limit <= 0 {
		limit = 10
	}

	posts, err := ph.postSrv.GetByUserID(user.ID, limit)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get posts")

		return
	}

	response.JSON(w, http.StatusOK, posts)
}

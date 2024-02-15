package handlers

import (
	"encoding/json"
	"github.com/makarellav/bloagg/internal/entity"
	"github.com/makarellav/bloagg/internal/http/response"
	"github.com/makarellav/bloagg/internal/services"
	"net/http"
)

type FeedHandlers struct {
	feedSrv *services.FeedService
}

func NewFeedHandlers(feedSrv *services.FeedService) *FeedHandlers {
	return &FeedHandlers{feedSrv: feedSrv}
}

func (fh *FeedHandlers) HandleCreate(w http.ResponseWriter, r *http.Request, user *entity.User) {
	var body entity.Feed

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to decode request body")

		return
	}

	body.UserID = user.ID
	newFeed, err := fh.feedSrv.Create(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create a feed")

		return
	}

	response.JSON(w, http.StatusCreated, newFeed)
}

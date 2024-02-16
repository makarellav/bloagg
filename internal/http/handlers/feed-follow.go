package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/makarellav/bloagg/internal/entity"
	"github.com/makarellav/bloagg/internal/http/response"
	"github.com/makarellav/bloagg/internal/services"
	"net/http"
)

type FeedFollowHandlers struct {
	feedFollowSrv *services.FeedFollowService
}

func NewFeedFollowHandlers(feedFollowSrv *services.FeedFollowService) *FeedFollowHandlers {
	return &FeedFollowHandlers{feedFollowSrv: feedFollowSrv}
}

func (ffh *FeedFollowHandlers) HandleCreate(w http.ResponseWriter, r *http.Request, user *entity.User) {
	var body entity.FeedFollow

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to decode request body")

		return
	}

	body.UserID = user.ID

	newFeedFollow, err := ffh.feedFollowSrv.Create(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create a feed follow")

		return
	}

	response.JSON(w, http.StatusCreated, newFeedFollow)
}

func (ffh *FeedFollowHandlers) HandleDelete(w http.ResponseWriter, r *http.Request, user *entity.User) {
	feedFollowIDParam := r.PathValue("feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDParam)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid feed follow id")

		return
	}

	feedFollow := entity.FeedFollow{
		ID:     feedFollowID,
		UserID: user.ID,
	}

	err = ffh.feedFollowSrv.Delete(feedFollow)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to delete a feed follow")

		return
	}

	response.JSON(w, http.StatusOK, struct{}{})
}

func (ffh *FeedFollowHandlers) HandleGetAllByUserID(w http.ResponseWriter, _ *http.Request, user *entity.User) {
	feedFollows, err := ffh.feedFollowSrv.GetAllByUserID(user.ID)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get follow feeds")

		return
	}

	response.JSON(w, http.StatusOK, feedFollows)
}

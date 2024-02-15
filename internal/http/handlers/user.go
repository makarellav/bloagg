package handlers

import (
	"encoding/json"
	"github.com/makarellav/bloagg/internal/entity"
	"github.com/makarellav/bloagg/internal/http/response"
	"github.com/makarellav/bloagg/internal/services"
	"net/http"
)

type UserHandlers struct {
	userSrv *services.UserService
}

func NewUserHandlers(userSrv *services.UserService) *UserHandlers {
	return &UserHandlers{userSrv: userSrv}
}

func (uh *UserHandlers) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var body entity.User

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to decode request body")

		return
	}

	newUser, err := uh.userSrv.Create(&body)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to create a user")

		return
	}

	response.JSON(w, http.StatusCreated, newUser)
}

func (uh *UserHandlers) HandleGetByKey(w http.ResponseWriter, _ *http.Request, user *entity.User) {
	response.JSON(w, http.StatusOK, user)
}

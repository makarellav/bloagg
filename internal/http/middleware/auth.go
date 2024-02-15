package middleware

import (
	"github.com/makarellav/bloagg/internal/entity"
	"github.com/makarellav/bloagg/internal/http/auth"
	"github.com/makarellav/bloagg/internal/http/response"
	"github.com/makarellav/bloagg/internal/services"
	"net/http"
)

type AuthMiddleware struct {
	userSrv *services.UserService
}

type userHandler func(w http.ResponseWriter, r *http.Request, user *entity.User)

func NewAuthMiddleware(userSrv *services.UserService) *AuthMiddleware {
	return &AuthMiddleware{userSrv: userSrv}
}

func (am *AuthMiddleware) Auth(h userHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			response.Error(w, http.StatusUnauthorized, err.Error())

			return
		}

		user, err := am.userSrv.GetByKey(apiKey)

		if err != nil {
			response.Error(w, http.StatusNotFound, "not found")

			return
		}

		h(w, r, user)
	}
}

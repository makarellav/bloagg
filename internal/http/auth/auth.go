package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrUnauthorized = errors.New("unauthorized")

func GetAPIKey(h http.Header) (string, error) {
	authHeader := h.Get("Authorization")

	if authHeader == "" {
		return "", ErrUnauthorized
	}

	authKey := strings.Split(authHeader, " ")

	if len(authKey) < 2 || authKey[0] != "ApiKey" {
		return "", ErrUnauthorized
	}

	return authKey[1], nil
}

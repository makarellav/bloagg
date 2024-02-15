package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("respondWithJSON: error marshalling json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(status)
	w.Write(data)
}

func Error(w http.ResponseWriter, status int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	JSON(w, status, errorResponse{
		Error: message,
	})
}

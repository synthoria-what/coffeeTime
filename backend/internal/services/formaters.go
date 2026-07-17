package services

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"synthori.space/coffeeTime/internal/models"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(data); err != nil {
		log.Printf("wrong json data: %s", err)
	}
}

func WriteError(w http.ResponseWriter, status int, message string) {

	now := time.Now().UTC()

	resp := models.ErrorResponse{
		Message:    message,
		Status:     status,
		StatusText: http.StatusText(status),
		Timestamp:  now,
	}

	WriteJSON(w, status, resp)
}

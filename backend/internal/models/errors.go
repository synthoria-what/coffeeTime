package models

import "time"

type ErrorResponse struct {
	Message    string    `json:"message"`
	Status     int       `json:"status"`
	StatusText string    `json:"status_text"`
	Timestamp  time.Time `json:"timestamp"`
}

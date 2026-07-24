package messages

import (
	"errors"
	"time"
)

var ErrUserNotFound = errors.New("user not found")
var ErrUserNotHavePermisions = errors.New("you dont have permisions")
var ErrDontHavePermisions = errors.New("you dont have permisions")

var ErrInvalidToken = errors.New("invalid token")

type RegisterResponse struct {
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	TimeStamp time.Time `json:"timestamp"`
}

package messages

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrUserNotHavePermisions = errors.New("you dont have permisions")
var ErrDontHavePermisions = errors.New("you dont have permisions")

var ErrInvalidToken = errors.New("invalid token")

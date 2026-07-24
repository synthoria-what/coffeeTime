package models

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserID int    `json:"id"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

type TokenResponse struct {
	Token string `json:"token"`
}

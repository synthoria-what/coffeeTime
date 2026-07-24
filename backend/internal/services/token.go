package services

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"synthori.space/coffeeTime/internal/models"
)

var jwtSecret = []byte("super-secret-key-123123123")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(hashPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashPassword, password)
}

func GetTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("authorization header is missing")
	}

	parts := strings.SplitN(header, " ", 2)

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header")
	}

	if parts[1] == "" {
		return "", errors.New("token is empty")
	}

	return parts[1], nil
}

func ParseToken(tokenString string) (*models.TokenClaims, error) {
	claims := &models.TokenClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return jwtSecret, nil
		},
		jwt.WithValidMethods(
			[]string{jwt.SigningMethodHS256.Alg()},
		),
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GenerateToken(data models.UserAuthData) (string, error) {
	now := time.Now()

	claims := models.TokenClaims{
		UserID: data.ID,
		Role:   data.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(
				now.Add(2 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(jwtSecret)
}

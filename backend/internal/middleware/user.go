package middleware

import (
	"context"
	"fmt"
	"net/http"

	"synthori.space/coffeeTime/internal/messages"
	"synthori.space/coffeeTime/internal/services"
)

type contextKey string

const (
	UserIDKey   contextKey = "userID"
	UserRoleKey contextKey = "userRole"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := services.GetTokenFromHeader(r)
		if err != nil {
			services.WriteError(w, http.StatusForbidden, messages.ErrInvalidToken.Error())
			return
		}

		claims, err := services.ParseToken(tokenString)
		if err != nil {
			services.WriteError(w, http.StatusForbidden, messages.ErrInvalidToken.Error())
			return
		}

		ctx := context.WithValue(
			r.Context(),
			UserIDKey,
			claims.UserID,
		)

		ctx = context.WithValue(
			ctx,
			UserRoleKey,
			claims.Role,
		)

		fmt.Println("ID пользователя:", claims.UserID)
		fmt.Println("Роль пользователя:", claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"synthori.space/coffeeTime/internal/messages"
	"synthori.space/coffeeTime/internal/middleware"
	"synthori.space/coffeeTime/internal/services"
	"synthori.space/coffeeTime/mock"
)

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/error", mock.SendErrorNotFound)
	r.Post("/generate_users", GenerateMockUsers)
	r.Post("/generate_database", GenerateMockDatabase)

	r.Route("/me", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			userID, ok := r.Context().Value(middleware.UserIDKey).(int)
			if !ok {
				services.WriteError(w, http.StatusUnauthorized, messages.ErrInvalidToken.Error())
				return
			}

			user, err := services.GetUser(userID)
			if err != nil {
				services.WriteError(w, http.StatusBadRequest, "user not found")
				return
			}

			services.WriteJSON(w, http.StatusOK, user)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", GetUsers)
		r.Get("/{id}", GetUser)
	})

	return r
}

func GenerateMockDatabase(w http.ResponseWriter, r *http.Request) {
	err := services.GenerateDatabase()
	if err != nil {
		services.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	services.WriteJSON(w, http.StatusAccepted, "genrate database")
}

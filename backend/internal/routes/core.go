package routes

import (
	"github.com/go-chi/chi/v5"
	"synthori.space/coffeeTime/mock"
)

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/error", mock.SendErrorNotFound)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", GetUser)
	})

	return r
}

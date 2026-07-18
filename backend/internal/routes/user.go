package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"synthori.space/coffeeTime/internal/database"
	"synthori.space/coffeeTime/internal/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	users, err := services.GetUsers(limit, offset)

	services.WriteJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		services.WriteError(w, http.StatusBadRequest, "invalid id")
	}

	user, err := services.GetUser(id)

	services.WriteJSON(w, http.StatusOK, user)
}

func GetMyProfile(w http.ResponseWriter, r *http.Request) {

}

func GenerateMockUsers(w http.ResponseWriter, r *http.Request) {
	database.GenerateRandomUsers(20)
	services.WriteJSON(w, http.StatusAccepted, nil)
}

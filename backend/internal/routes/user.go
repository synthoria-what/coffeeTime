package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"synthori.space/coffeeTime/internal/database"
	"synthori.space/coffeeTime/internal/messages"
	"synthori.space/coffeeTime/internal/middleware"
	"synthori.space/coffeeTime/internal/models"
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

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		services.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := services.Login(req)
	if err != nil {
		services.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	response := models.TokenResponse{
		Token: token,
	}

	services.WriteJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		services.WriteError(w, http.StatusBadRequest, err.Error())
	}

	_, err := services.Registeruser(req)
	if err != nil {
		services.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now().UTC()

	resp := messages.RegisterResponse{
		Message:   http.StatusText(http.StatusOK),
		Status:    http.StatusOK,
		TimeStamp: now,
	}

	services.WriteJSON(w, http.StatusOK, resp)
}

func GetMyProfile(w http.ResponseWriter, r *http.Request) {
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
}

func GenerateMockUsers(w http.ResponseWriter, r *http.Request) {
	database.GenerateRandomUsers(20)
	services.WriteJSON(w, http.StatusAccepted, nil)
}

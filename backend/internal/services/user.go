package services

import (
	"errors"

	"synthori.space/coffeeTime/internal/database"
	"synthori.space/coffeeTime/internal/models"
)

func GetUsers(limit int, offset int) ([]models.GetUserResponse, error) {
	users, err := database.GetUsers(limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(userID int) (models.GetUserResponse, error) {
	user, err := database.GetUserByID(userID)
	if err != nil {
		return models.GetUserResponse{}, err
	}

	return user, nil
}

func CreateUser(req models.CreateUserRequest) (models.User, error) {
	exists, err := database.UserExistsByUsername(req.Username)
	if err != nil {
		return models.User{}, err
	}

	if exists {
		return models.User{}, errors.New("username already exists")
	}

	passwordHash, err := HashPassword(req.Password)
	if err != nil {
		return models.User{}, err
	}

	user := models.CreateUserRequest{
		Username:   req.Username,
		Password:   passwordHash,
		Role:       req.Role,
		AvatarPath: req.AvatarPath,
	}

	createdUser, err := database.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

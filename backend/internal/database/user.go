package database

import (
	"fmt"
	"math/rand"
	"time"

	"synthori.space/coffeeTime/internal/messages"
	"synthori.space/coffeeTime/internal/models"
)

func GenerateRandomUsers(size int) error {

	if size < 0 {
		size = 10
	}

	usernames := []string{
		"Ivan",
		"Nikita",
		"synthoria",
		"bobster",
		"Sigma123124",
		"kirill",
		"roman",
		"nikita",
		"andrey",
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		randomIndex := rng.Intn(len(usernames))
		randomName := usernames[randomIndex]

		// Суффикс нужен, чтобы username не повторялся в базе.
		username := fmt.Sprintf("%s_%d", randomName, i+1)

		req := models.CreateUserRequest{
			Username:   username,
			Password:   username,
			Role:       "user",
			AvatarPath: "",
		}

		CreateUser(req)
	}

	return nil

}

func GetUserByID(userID int) (models.GetUserResponse, error) {
	var user models.GetUserResponse

	err := db.QueryRow("select id, username, role, avatar_path from users where id=?", userID).Scan(
		&user.ID,
		&user.Username,
		&user.Role,
		&user.AvatarPath,
	)

	if err != nil {
		return models.GetUserResponse{}, messages.ErrUserNotFound
	}

	return user, nil

}

func GetFullUserProfile(userID int) (models.User, error) {
	var user models.User

	err := db.QueryRow("select id, username, password, role, avatar_path from users where id=?", userID).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.Role,
		&user.AvatarPath,
	)

	if err != nil {
		return models.User{}, messages.ErrDontHavePermisions
	}

	return user, nil
}

func UserExistsByUsername(username string) (bool, error) {
	var count int

	err := db.QueryRow("select count(*) from users where username=?", username).Scan(
		&count,
	)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetUsers(limit int, offset int) ([]models.GetUserResponse, error) {
	users := []models.GetUserResponse{}

	rows, err := db.Query("select id, username, role, avatar_path from users limit ? offset ?", limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.GetUserResponse
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Role,
			&user.AvatarPath,
		)

		if err != nil {
			return nil, err
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}
		defer rows.Close()

		users = append(users, user)
	}

	return users, nil
}

func CreateUser(req models.CreateUserRequest) (models.User, error) {
	result, err := db.Exec("insert into users (username, password, role, avatar_path) values (?, ?, ?, ?)", req.Username, req.Password, req.Role, req.AvatarPath)
	if err != nil {
		return models.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		ID:           int(id),
		Username:     req.Username,
		PasswordHash: req.Password,
		Role:         req.Role,
		AvatarPath:   req.AvatarPath,
	}

	return user, nil
}

package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	PasswordHash string `json:"password_hash"`
	AvatarPath   string `json:"avatar_path"`
}

type GetUserResponse struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	AvatarPath string `json:"avatar_path"`
}

type CreateUserRequest struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	Password   string `json:"password"`
	AvatarPath string `json:"avatar_path"`
}

type UpdateUserRequest struct {
	Username   *string `json:"username"`
	Role       *string `json:"role"`
	AvatarPath *string `json:"avatar_path"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

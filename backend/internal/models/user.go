package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Role         string `json:"role"`
	PasswordHash string `json:"-"`
	AvatarPath   string `json:"avatar_path"`
}

type GetUserResponse struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	AvatarPath string `json:"avatar_path"`
}

type UserAuthData struct {
	ID           int
	PasswordHash string
	Role         string
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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

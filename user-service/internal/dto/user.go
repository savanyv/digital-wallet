package dtos

type CreateUserRequest struct {
	UserID string  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserResponse struct {
	UserID string  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

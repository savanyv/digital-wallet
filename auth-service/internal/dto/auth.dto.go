package dtos

type RegisterRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	UserId string `json:"user_id"`
	Token string `json:"token"`
	Message string `json:"message"`
}

package usecase

import (
	"errors"

	dtos "github.com/savanyv/digital-wallet/auth-service/internal/dto"
	"github.com/savanyv/digital-wallet/auth-service/internal/models"
	"github.com/savanyv/digital-wallet/auth-service/internal/repository"
	"github.com/savanyv/digital-wallet/shared/utils/bcrypt"
	"github.com/savanyv/digital-wallet/shared/utils/jwt"
)

type AuthUsecase interface {
	Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error)
	Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error)
}

type authUsecase struct {
	repo repository.AuthRepository
	jwt jwt.JWTService
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &authUsecase{
		repo: repo,
		jwt: jwt.NewJWTService(),
	}
}

func (u *authUsecase) Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error) {
	// check if email exists
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("Email already exists")
	}

	// hash password
	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("Error hashing password")
	}

	// create user
	user = &models.Auth{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := u.repo.Create(user); err != nil {
		return nil, errors.New("Error creating user")
	}

	resp := &dtos.AuthResponse{
		UserId: string(user.ID),
		Message: "User created successfully",
	}

	return resp, nil
}

func (u *authUsecase) Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error) {
	// check if email exists
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("User not found")
	}

	// compare password
	if err := bcrypt.ComparePassword(user.Password, req.Password); err != nil {
		return nil, errors.New("Invalid email or password")
	}

	// generate token
	token, err := u.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("Error generating token")
	}

	resp := &dtos.AuthResponse{
		UserId: string(user.ID),
		Token: token,
		Message: "Login successful",
	}

	return resp, nil
}

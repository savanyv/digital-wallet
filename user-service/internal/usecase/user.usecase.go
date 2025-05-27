package usecase

import (
	"errors"

	"github.com/google/uuid"
	dtos "github.com/savanyv/digital-wallet/user-service/internal/dto"
	"github.com/savanyv/digital-wallet/user-service/internal/models"
	"github.com/savanyv/digital-wallet/user-service/internal/repository"
)

type UserUsecase interface {
	CreateUser(req *dtos.CreateUserRequest) (*dtos.UserResponse, error)
	FindUserByID(ID string) (*dtos.UserResponse, error)
	FindUserByEmail(email string) (*dtos.UserResponse, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecsae(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) CreateUser(req *dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	existing, err := u.repo.FindUserByID(req.UserID)
	if err == nil && existing != nil {
		return nil, errors.New("user already exists")
	}

	user := &models.User{
		ID:    userUUID,
		Name:  req.Name,
		Email: req.Email,
	}

	if err := u.repo.Create(user); err != nil {
		return nil, errors.New("error creating user")
	}

	resp := &dtos.UserResponse{
		UserID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}

	return resp, nil
}

func (u *userUsecase) FindUserByID(ID string) (*dtos.UserResponse, error) {
	user, err := u.repo.FindUserByID(ID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := &dtos.UserResponse{
		UserID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}

	return resp, nil
}

func (u *userUsecase) FindUserByEmail(email string) (*dtos.UserResponse, error) {
	user, err := u.repo.FindUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	resp := &dtos.UserResponse{
		UserID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}

	return resp, nil
}

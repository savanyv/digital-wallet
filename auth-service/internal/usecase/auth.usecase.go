package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/savanyv/digital-wallet/auth-service/internal/client"
	dtos "github.com/savanyv/digital-wallet/auth-service/internal/dto"
	"github.com/savanyv/digital-wallet/auth-service/internal/models"
	"github.com/savanyv/digital-wallet/auth-service/internal/repository"
	userPB "github.com/savanyv/digital-wallet/proto/user"
	"github.com/savanyv/digital-wallet/shared/utils/bcrypt"
	"github.com/savanyv/digital-wallet/shared/utils/jwt"
	walletPB "github.com/savanyv/digital-wallet/proto/wallet"
)

type AuthUsecase interface {
	Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error)
	Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error)
}

type authUsecase struct {
	repo repository.AuthRepository
	jwt jwt.JWTService
	userClient client.UserGrpcClient
	walletClient client.WalletGrpcClient
}

func NewAuthUsecase(repo repository.AuthRepository, userClient client.UserGrpcClient, walletClient client.WalletGrpcClient) AuthUsecase {
	return &authUsecase{
		repo: repo,
		jwt: jwt.NewJWTService(),
		userClient: userClient,
		walletClient: walletClient,
	}
}

func (u *authUsecase) Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error) {
	// check if email exists
	user, err := u.repo.FindUserByEmail(req.Email)
	if err == nil && user != nil {
		return nil, errors.New("email already exists")
	}

	// hash password
	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("error hashing password")
	}

	// create user
	user = &models.Auth{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := u.repo.Create(user); err != nil {
		return nil, errors.New("error creating user")
	}

	_, err = u.userClient.CreateUser(context.Background(), &userPB.CreateUserRequest{
		UserId: user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
	})
	if err != nil {
		log.Println("Warning: failed to sync with User-Service: ", err)
	}

	_, err = u.walletClient.CreateWallet(context.Background(), &walletPB.CreateWalletRequest{
		UserId: user.ID.String(),
	})
	if err != nil {
		log.Println("Warning: failed to sync with Wallet-Service: ", err)
	}

	resp := &dtos.AuthResponse{
		UserId: user.ID.String(),
		Message: "User created successfully",
	}

	return resp, nil
}

func (u *authUsecase) Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error) {
	// check if email exists
	user, err := u.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// compare password
	if err := bcrypt.ComparePassword(user.Password, req.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// generate token
	token, err := u.jwt.GenerateToken(user.ID.String(), user.Email)
	if err != nil {
		return nil, errors.New("error generating token")
	}

	resp := &dtos.AuthResponse{
		UserId: user.ID.String(),
		Token: token,
		Message: "Login successful",
	}

	return resp, nil
}

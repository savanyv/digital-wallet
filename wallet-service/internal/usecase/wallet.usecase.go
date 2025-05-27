package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	dtos "github.com/savanyv/digital-wallet/wallet-service/internal/dto"
	"github.com/savanyv/digital-wallet/wallet-service/internal/models"
	"github.com/savanyv/digital-wallet/wallet-service/internal/repository"
)

type WalletUsecase interface {
	CreateWallet(req *dtos.CreateWalletRequest) (*dtos.WalletResponse, error)
	GetWallet(req *dtos.GetWalletRequest) (*dtos.WalletResponse, error)
	UpdateBalance(req *dtos.UpdateBalanceRequest) (*dtos.WalletResponse, error)
}

type walletUsecase struct {
	repo repository.WalletRepository
}

func NewWalletUsecase(repo repository.WalletRepository) WalletUsecase {
	return &walletUsecase{
		repo: repo,
	}
}

func (u *walletUsecase) CreateWallet(req *dtos.CreateWalletRequest) (*dtos.WalletResponse, error) {
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	wallet, err := u.repo.GetUserByID(req.UserID)
	if err == nil && wallet != nil {
		return nil, errors.New("wallet already exists for this user")
	}

	newWallet := &models.Wallet{
		UserID: userUUID,
		Balance: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = u.repo.Create(newWallet)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}

	resp := &dtos.WalletResponse{
		UserID: newWallet.UserID.String(),
		Balance: newWallet.Balance,
		Message: "wallet created successfully",
		CreatedAt: newWallet.CreatedAt,
		UpdatedAt: newWallet.UpdatedAt,
	}

	return resp, nil
}

func (u *walletUsecase) GetWallet(req *dtos.GetWalletRequest) (*dtos.WalletResponse, error) {
	wallet, err := u.repo.GetUserByID(req.UserID)
	if err != nil {
		return nil, err
	}

	resp := &dtos.WalletResponse{
		UserID: wallet.UserID.String(),
		Balance: wallet.Balance,
		Message: "wallet retrieved successfully",
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}

	return resp, nil
}

func (u *walletUsecase) UpdateBalance(req *dtos.UpdateBalanceRequest) (*dtos.WalletResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	wallet, err := u.repo.GetUserByID(req.UserID)
	if err != nil {
		return nil, errors.New("wallet not found")
	}

	const (
		operationDeposit = "deposit"
		operationWithdraw = "withdraw"
	)

	switch req.Operation {
		case operationDeposit:
			wallet.Balance += req.Amount
		case operationWithdraw:
			if wallet.Balance < req.Amount {
				return nil, errors.New("insufficient balance")
			}
			wallet.Balance -= req.Amount
		default:
			return nil, errors.New("invalid operation")
	}

	wallet.UpdatedAt = time.Now()
	err = u.repo.Update(wallet)
	if err != nil {
		return nil, errors.New("failed to update balance")
	}

	resp := &dtos.WalletResponse{
		UserID: wallet.UserID.String(),
		Balance: wallet.Balance,
		Message: "Balance updated successfully",
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}

	return resp, nil
}

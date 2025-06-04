package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	walletPB "github.com/savanyv/digital-wallet/proto/wallet"
	"github.com/savanyv/digital-wallet/transaction-service/internal/client"
	dtos "github.com/savanyv/digital-wallet/transaction-service/internal/dto"
	"github.com/savanyv/digital-wallet/transaction-service/internal/models"
	"github.com/savanyv/digital-wallet/transaction-service/internal/repository"
)

type TransactionUsecase interface {
	Deposit(req *dtos.DepositWithdrawRequest) (*dtos.TransactionResponse, error)
	Withdraw(req *dtos.DepositWithdrawRequest) (*dtos.TransactionResponse, error)
	Transfer(req *dtos.TransferRequest) (*dtos.TransactionResponse, error)
	GetHistory(userID string) ([]dtos.Transaction, error)
}

type transactionUsecase struct {
	repo repository.TransactionRepository
	walletClient client.WalletGrpcClient
}

func NewTransactionUsecase(repo repository.TransactionRepository, walletClient client.WalletGrpcClient) TransactionUsecase {
	return &transactionUsecase{
		repo: repo,
		walletClient: walletClient,
	}
}

func (u *transactionUsecase) Deposit(req *dtos.DepositWithdrawRequest) (*dtos.TransactionResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount mus be greater than zero")
	}

	_, err := u.walletClient.UpdateBalance(context.Background(), &walletPB.UpdateBalanceRequest{
		UserId: req.UserID,
		Amount: req.Amount,
		Operation: "deposit",
	})
	if err != nil {
		log.Println("WalletService error:", err)
		return nil, errors.New("failed to update wallet balance")
	}

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	t := &models.Transaction{
		UserID: userUUID,
		Type: "deposit",
		Amount: int64(req.Amount),
		Description: "Deposit To Wallet",
		CreatedAt: time.Now(),
	}

	if err := u.repo.Save(t); err != nil {
		return nil, errors.New("failed to save transaction")
	}

	resp := &dtos.TransactionResponse{
		UserID: t.UserID.String(),
		Type: t.Type,
		Amount: t.Amount,
		Message: "deposit successful",
		CreatedAt: t.CreatedAt,
	}

	return resp, nil
}

func (u *transactionUsecase) Withdraw(req *dtos.DepositWithdrawRequest) (*dtos.TransactionResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	_, err := u.walletClient.UpdateBalance(context.Background(), &walletPB.UpdateBalanceRequest{
		UserId: req.UserID,
		Amount: req.Amount,
		Operation: "withdraw",
	})
	if err != nil {
		log.Println("WalletService error:", err)
		return nil, errors.New("failed to update wallet balance")
	}

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	t := &models.Transaction{
		ID: uuid.New(),
		UserID: userUUID,
		Type: "withdraw",
		Amount: req.Amount,
		Description: "Withdraw From Wallet",
		CreatedAt: time.Now(),
	}

	resp := &dtos.TransactionResponse{
		UserID: t.UserID.String(),
		Type: t.Type,
		Amount: t.Amount,
		Message: "withdraw successful",
		CreatedAt: t.CreatedAt,
	}

	return resp, nil
}

func (u *transactionUsecase) Transfer(req *dtos.TransferRequest) (*dtos.TransactionResponse, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if req.SenderID == req.ReceiverID {
		return nil, errors.New("cannot transfer to the same user")
	}

	_, err := u.walletClient.UpdateBalance(context.Background(), &walletPB.UpdateBalanceRequest{
		UserId: req.SenderID,
		Amount: req.Amount,
		Operation: "withdraw",
	})

	if err != nil {
		log.Println("Withdraw error:", err)
		return nil, errors.New("failed to update wallet balance")
	}

	_, err = u.walletClient.UpdateBalance(context.Background(), &walletPB.UpdateBalanceRequest{
		UserId: req.ReceiverID,
		Amount: req.Amount,
		Operation: "deposit",
	})

	if err != nil {
		_, _ = u.walletClient.UpdateBalance(context.Background(), &walletPB.UpdateBalanceRequest{
			UserId: req.SenderID,
			Amount: req.Amount,
			Operation: "deposit",
		})
		log.Println("Deposit error:", err)
		return nil, errors.New("failed to deposit to receiver")
	}

	senderUUID, err := uuid.Parse(req.SenderID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	t := &models.Transaction{
		ID: uuid.New(),
		UserID: senderUUID,
		Type: "transfer",
		Amount: req.Amount,
		Description: "Transfer To " + req.ReceiverID,
		CreatedAt: time.Now(),
	}

	if err := u.repo.Save(t); err != nil {
		return nil, errors.New("failed to save transaction")
	}

	resp := &dtos.TransactionResponse{
		UserID: t.UserID.String(),
		Type: t.Type,
		Amount: t.Amount,
		Message: "transfer successful",
		CreatedAt: t.CreatedAt,
	}

	return resp, nil
}

func (u *transactionUsecase) GetHistory(userID string) ([]dtos.Transaction, error) {
	transactions, err := u.repo.GetByUserID(uuid.MustParse(userID))
	if err != nil {
		return nil, err
	}

	var results []dtos.Transaction
	for _, t := range transactions {
		results = append(results, dtos.Transaction{
			TransactionID: t.ID.String(),
			UserID: t.UserID.String(),
			Type: t.Type,
			Amount: t.Amount,
			Description: t.Description,
			CreatedAt: t.CreatedAt,
		})
	}

	return results, nil
}

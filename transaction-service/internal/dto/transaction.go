package dtos

import "time"

type DepositWithdrawRequest struct {
	UserID string `json:"user_id"`
	Amount int64  `json:"amount"`
}

type TransferRequest struct {
	SenderID string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Amount int64 `json:"amount"`
}

type TransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	UserID string `json:"user_id"`
	Type string `json:"type"`
	Amount int64 `json:"amount"`
	Message string `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type GetHistoryRequest struct {
	UserID string `json:"user_id"`
}

type Transaction struct {
	TransactionID string `json:"transaction_id"`
	UserID string `json:"user_id"`
	Type string `json:"type"`
	Amount int64 `json:"amount"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionHistoryResponse struct {
	Transactions []Transaction `json:"transactions"`
}

package dtos

import "time"

type CreateWalletRequest struct {
	UserID string `json:"user_id"`
}

type GetWalletRequest struct {
	UserID string `json:"user_id"`
}

type UpdateBalanceRequest struct {
	UserID string  `json:"user_id"`
	Amount int64 `json:"amount"`
	Operation string `json:"operation"`
}

type WalletResponse struct {
	UserID string  `json:"user_id"`
	Balance int64 `json:"balance"`
	Message string `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

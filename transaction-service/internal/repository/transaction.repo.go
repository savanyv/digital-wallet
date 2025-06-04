package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/savanyv/digital-wallet/transaction-service/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(tx *models.Transaction) error
	GetByUserID(userID uuid.UUID) ([]*models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) Save(tx *models.Transaction) error {
	if err := r.db.Create(tx).Error; err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) GetByUserID(userID uuid.UUID) ([]*models.Transaction, error) {
	var txs []*models.Transaction
	if err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&txs).Error; err != nil {
		return nil, err
	}
	
	if len(txs) == 0 {
		return nil, errors.New("no transactions found")
	}

	return txs, nil
}

package repository

import (
	"github.com/savanyv/digital-wallet/wallet-service/internal/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Create(wallet *models.Wallet) error
	GetUserByID(userID string) (*models.Wallet, error)
	Update(wallet *models.Wallet) error
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{
		db: db,
	}
}

func (r *walletRepository) Create(wallet *models.Wallet) error {
	if err := r.db.Create(wallet).Error; err != nil {
		return err
	}

	return nil
}

func (r *walletRepository) GetUserByID(userID string) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *walletRepository) Update(wallet *models.Wallet) error {
	if err := r.db.Save(wallet).Error; err != nil {
		return err
	}

	return nil
}

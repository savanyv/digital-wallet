package repository

import (
	"github.com/savanyv/digital-wallet/auth-service/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(req *models.Auth) error
	FindUserByEmail(email string) (*models.Auth, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Create(req *models.Auth) error {
	if err := r.db.Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepository) FindUserByEmail(email string) (*models.Auth, error) {
	var user models.Auth
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

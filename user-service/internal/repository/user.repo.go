package repository

import (
	"github.com/savanyv/digital-wallet/user-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(req *models.User) error
	FindUserByID(id string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(req *models.User) error {
	if err := r.db.Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

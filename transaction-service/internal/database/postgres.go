package database

import (
	"fmt"
	"log"

	"github.com/savanyv/digital-wallet/shared/config"
	"github.com/savanyv/digital-wallet/transaction-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBTHost,
		cfg.DBTUser,
		cfg.DBTPass,
		cfg.DBTName,
		cfg.DBTPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database")
		return nil, err
	}

	if err := db.AutoMigrate(&models.Transaction{}); err != nil {
		log.Println("Failed to migrate database")
		return nil, err
	}

	DB = db
	return db, nil
}

package database

import (
	"fmt"
	"log"

	"github.com/savanyv/digital-wallet/shared/config"
	"github.com/savanyv/digital-wallet/wallet-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBWHost,
		cfg.DBWUser,
		cfg.DBWPass,
		cfg.DBWName,
		cfg.DBWPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}

	if err := db.AutoMigrate(&models.Wallet{}); err != nil {
		log.Println("Error migrating database")
		return nil, err
	}

	DB = db
	return db, nil
}

package database

import (
	"fmt"
	"log"

	"github.com/savanyv/digital-wallet/auth-service/internal/models"
	"github.com/savanyv/digital-wallet/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPosgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBAHost,
		cfg.DBAUser,
		cfg.DBAPassword,
		cfg.DBAHost,
		cfg.DBAPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}

	if err := db.AutoMigrate(&models.Auth{}); err != nil {
		log.Println("Error migrating database")
		return nil, err
	}

	DB = db
	return db, nil
}

package database

import (
	"fmt"
	"log"

	"github.com/savanyv/digital-wallet/shared/config"
	"github.com/savanyv/digital-wallet/user-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBUHost,
		cfg.DBUUser,
		cfg.DBUPass,
		cfg.DBUName,
		cfg.DBUPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Println("Error migrating database")
		return nil, err
	}

	DB = db
	return db, nil
}

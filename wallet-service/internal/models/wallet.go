package models

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID    int64  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserID uuid.UUID  `json:"user_id" gorm:"type:uuid;gen_random_uuid;not null"`
	Balance float64 `json:"balance" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

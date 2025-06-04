package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid();not null"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Type        string    `json:"type" gorm:"type:varchar(20);not null"`
	Amount      int64     `json:"amount" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time
}

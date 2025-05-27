package models

import "github.com/google/uuid"

type Auth struct {
	ID   uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid();not null"`
	Name string `json:"name" gorm:"type:varchar(255); not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

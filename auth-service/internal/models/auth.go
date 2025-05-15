package models

type Auth struct {
	ID   int64  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name" gorm:"type:varchar(255); not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}

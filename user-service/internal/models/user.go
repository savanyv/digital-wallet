package models

type User struct {
	ID int64  `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name" gorm:"type:varchar(255); not null"`
	Email string `json:"email" gorm:"type:varchar(255);not null"`
}

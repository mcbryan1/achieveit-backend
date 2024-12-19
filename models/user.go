package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username   string    `json:"username" gorm:"type:varchar(100);not null"`
	Password   string    `json:"password" gorm:"type:varchar(100);not null"`
	Goals      []Goal    `json:"goals" gorm:"foreignKey:UserID;references:ID"`
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Goal struct {
	gorm.Model  `json:"-"`
	ID          uuid.UUID   `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string      `json:"title" gorm:"type:text;not null"`
	Description string      `json:"description" gorm:"type:text;not null"`
	Progress    float64     `json:"progress" gorm:"type:float;not null"`
	UserID      uuid.UUID   `json:"user_id" gorm:"type:uuid;not null;references:ID"`
	Milestones  []Milestone `json:"milestones" gorm:"foreignKey:GoalID;references:ID"`
}

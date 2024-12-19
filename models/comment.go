package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model  `json:"-"`
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	MilestoneID uuid.UUID `json:"milestone_id" gorm:"type:uuid;not null;references:ID"`
}

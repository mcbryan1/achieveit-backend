package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Milestone struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title      string    `json:"title" gorm:"type:text;not null"`
	Completed  bool      `json:"completed" gorm:"type:boolean;not null"`
	GoalID     uuid.UUID `json:"goal_id" gorm:"type:uuid;not null;references:ID"`
	Comments   []Comment `json:"comments" gorm:"foreignKey:MilestoneID;references:ID"`
}

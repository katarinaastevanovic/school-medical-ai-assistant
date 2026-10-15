package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	UserID       uint   `gorm:"not null"`
	WorkingHours string `json:"workingHours"`
}

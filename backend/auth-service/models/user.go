package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	UMCN      string `gorm:"unique;not null"`
	FirstName string `gorm:"size:100; null"`
	LastName  string `gorm:"size:100; null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      Role   `gorm:"type:varchar(20);not null"`
}

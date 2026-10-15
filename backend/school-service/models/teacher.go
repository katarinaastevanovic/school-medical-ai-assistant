package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model

	UserID uint `gorm:"unique;not null"`

	ConsultationHours string `json:"consultationHours"`

	Subjects []Subject `gorm:"many2many:teacher_subjects;"`
}

package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	UserID uint `gorm:"unique;not null"`

	ClassID uint
	Class   SchoolClass

	Grades   []Grade
	Absences []Absence
}

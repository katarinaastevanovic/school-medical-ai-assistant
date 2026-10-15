package models

import (
	"time"

	"gorm.io/gorm"
)

type Absence struct {
	gorm.Model

	StudentID uint
	Student   Student

	SubjectID uint
	Subject   Subject

	Status AbsenceType

	Date time.Time
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model

	Value int
	Date  time.Time

	StudentID uint
	Student   Student

	TeacherID uint
	Teacher   Teacher

	SubjectID uint
	Subject   Subject
}

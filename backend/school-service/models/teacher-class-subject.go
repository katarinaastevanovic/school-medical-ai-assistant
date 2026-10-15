package models

import "gorm.io/gorm"

type TeacherClassSubject struct {
	gorm.Model

	TeacherID uint
	Teacher   Teacher

	ClassID uint
	Class   SchoolClass

	SubjectID uint
	Subject   Subject
}

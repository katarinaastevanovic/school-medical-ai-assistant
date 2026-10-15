package models

import "gorm.io/gorm"

type Subject struct {
	gorm.Model

	Name string `gorm:"unique"`

	Grades []Grade

	Teachers []Teacher `gorm:"many2many:teacher_subjects;"`

	Classes []SchoolClass `gorm:"many2many:class_subjects;"`
}

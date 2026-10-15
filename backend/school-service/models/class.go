package models

import "gorm.io/gorm"

type SchoolClass struct {
	gorm.Model

	Name string

	Year int

	Students []Student `gorm:"foreignKey:ClassID"`

	Subjects []Subject `gorm:"many2many:class_subjects;"`
}

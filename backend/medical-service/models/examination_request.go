package models

import "gorm.io/gorm"

type ExaminationRequest struct {
	gorm.Model
	PatientID       uint              `gorm:"not null"`
	DoctorID        uint              `gorm:"not null"`
	Type            TypeOfExamination `gorm:"type:varchar(20);not null"`
	Status          TypeOfRequest     `gorm:"type:varchar(20);not null"`
	MedicalRecordID uint
	Examinations    []Examination `gorm:"foreignKey:RequestID"`
}

type TypeOfRequest string

const (
	REQUESTED TypeOfRequest = "REQUESTED"
	APPROVED  TypeOfRequest = "APPROVED"
	REJECTED  TypeOfRequest = "REJECTED"
)

type TypeOfExamination string

const (
	REGULAR    TypeOfExamination = "REGULAR"
	SPECIALIST TypeOfExamination = "SPECIALIST"
	URGENT     TypeOfExamination = "URGENT"
)

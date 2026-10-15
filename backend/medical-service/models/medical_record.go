package models

import (
	"time"

	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	PatientID       uint   `gorm:"not null"`
	Allergies       string `gorm:"type:text"`
	ChronicDiseases string `gorm:"type:text"`
	LastUpdate      time.Time
	Examinations    []Examination        `gorm:"foreignKey:MedicalRecordID"`
	Requests        []ExaminationRequest `gorm:"foreignKey:MedicalRecordID"`
}

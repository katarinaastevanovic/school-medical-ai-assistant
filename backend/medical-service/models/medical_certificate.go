package models

import (
	"time"

	"gorm.io/gorm"
)

type MedicalCertificate struct {
	gorm.Model
	PatientID uint              `gorm:"not null"`
	DoctorID  uint              `gorm:"not null"`
	Date      time.Time         `gorm:"not null"`
	Type      TypeOfCertificate `gorm:"type:varchar(20);not null"`
	Note      string            `gorm:"type:text"`
}

type TypeOfCertificate string

const (
	CERT_REGULAR   TypeOfCertificate = "REGULAR"
	CERT_PE        TypeOfCertificate = "PE"
	CERT_EXCURSION TypeOfCertificate = "EXCURSION"
	CERT_SICKNESS  TypeOfCertificate = "SICKNESS"
)

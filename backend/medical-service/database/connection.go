package database

import (
	"log"
	"os"

	"medical-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(
		&models.Doctor{},
		&models.Patient{},
		&models.MedicalRecord{},
		&models.ExaminationRequest{},
		&models.Examination{},
		&models.MedicalCertificate{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Medical database connected")
}

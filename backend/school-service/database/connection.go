package database

import (
	"log"
	"os"

	"school-service/models"

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
		&models.Subject{},
		&models.Teacher{},
		&models.SchoolClass{},
		&models.Student{},
		&models.Absence{},
		&models.Grade{},
		&models.TeacherClassSubject{},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}

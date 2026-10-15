package database

import (
	"chatbot-service/models"
	"log"
	"os"

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

	err = DB.AutoMigrate(&models.BotInstruction{}, &models.BotSettings{})
	if err != nil {
		log.Fatal(err)
	}
	var count int64

	DB.Model(&models.BotSettings{}).Count(&count)
	if count == 0 {
		DB.Create(&models.BotSettings{Temperature: 0.7, MaxTokens: 512, GroqModel: "openai/gpt-oss-120b"})
	}

	log.Println("Chatbot database connected")
}

package repository

import (
	"chatbot-service/database"
	"chatbot-service/models"
)

type ChatbotRepository interface {
	GetActiveInstructions(role string) ([]models.BotInstruction, error)
	GetAllInstructions() ([]models.BotInstruction, error)
	CreateInstruction(i *models.BotInstruction) error
	UpdateInstruction(i *models.BotInstruction) error
	DeleteInstruction(id uint) error
	GetSettings() (models.BotSettings, error)
	UpdateSettings(s *models.BotSettings) error
}

type chatbotRepository struct{}

func NewChatbotRepository() ChatbotRepository {
	return &chatbotRepository{}
}

func (r *chatbotRepository) GetActiveInstructions(role string) ([]models.BotInstruction, error) {
	var instructions []models.BotInstruction
	err := database.DB.
		Where("enabled = ? AND (role = ? OR role = '')", true, role).
		Order("\"order\" asc").
		Find(&instructions).Error
	return instructions, err
}

func (r *chatbotRepository) GetAllInstructions() ([]models.BotInstruction, error) {
	var instructions []models.BotInstruction
	err := database.DB.Order("\"order\" asc").Find(&instructions).Error
	return instructions, err
}

func (r *chatbotRepository) CreateInstruction(i *models.BotInstruction) error {
	return database.DB.Create(i).Error
}

func (r *chatbotRepository) UpdateInstruction(i *models.BotInstruction) error {
	return database.DB.Model(&models.BotInstruction{}).Where("id = ?", i.ID).Updates(i).Error
}

func (r *chatbotRepository) DeleteInstruction(id uint) error {
	return database.DB.Delete(&models.BotInstruction{}, id).Error
}

func (r *chatbotRepository) GetSettings() (models.BotSettings, error) {
	var s models.BotSettings
	err := database.DB.First(&s).Error
	return s, err
}

func (r *chatbotRepository) UpdateSettings(s *models.BotSettings) error {
	return database.DB.Model(&models.BotSettings{}).Where("id = ?", s.ID).Updates(s).Error
}

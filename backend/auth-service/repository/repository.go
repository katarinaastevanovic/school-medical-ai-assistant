package repository

import (
	"auth-service/database"
	"auth-service/models"
)

type AuthRepository interface {
	Create(user *models.Member) error
	FindByEmail(email string) (*models.Member, error)
	FindByUMCN(umcn string) (*models.Member, error)
	FindByID(id uint) (*models.Member, error)
	GetAllMembers() ([]models.Member, error)
	GetAllStudents() ([]models.Member, error)
	GetAllTeachers() ([]models.Member, error)
	GetAllDoctors() ([]models.Member, error)
	UpdateMember(member *models.Member) error
}

func (r *authRepository) FindByUMCN(umcn string) (*models.Member, error) {

	var user models.Member

	err := database.DB.
		Where("umcn = ?", umcn).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

type authRepository struct{}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) Create(user *models.Member) error {
	return database.DB.Create(user).Error
}

func (r *authRepository) FindByEmail(email string) (*models.Member, error) {
	var user models.Member

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) FindByID(id uint) (*models.Member, error) {
	var user models.Member

	err := database.DB.
		First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetAllMembers() ([]models.Member, error) {
	var members []models.Member
	result := database.DB.Find(&members)
	return members, result.Error
}

func (r *authRepository) GetAllStudents() ([]models.Member, error) {
	var members []models.Member
	result := database.DB.Where("role = ?", "STUDENT").Find(&members)
	return members, result.Error
}

func (r *authRepository) GetAllTeachers() ([]models.Member, error) {
	var members []models.Member
	result := database.DB.Where("role = ?", "TEACHER").Find(&members)
	return members, result.Error
}

func (r *authRepository) GetAllDoctors() ([]models.Member, error) {
	var members []models.Member
	result := database.DB.Where("role = ?", "DOCTOR").Find(&members)
	return members, result.Error
}

func (r *authRepository) UpdateMember(member *models.Member) error {
	return database.DB.Save(member).Error
}

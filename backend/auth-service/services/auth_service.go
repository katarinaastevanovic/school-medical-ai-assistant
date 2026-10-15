package services

import (
	"auth-service/dto"
	"auth-service/models"
	"auth-service/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo repository.AuthRepository
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *AuthService) Register(user *models.Member) error {

	if !ValidateJMBG(user.UMCN) {
		return errors.New("invalid jmbg")
	}

	if !ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}

	if user.Role == "" {
		return errors.New("role is required")
	}

	if user.FirstName == "" {
		return errors.New("first name is required")
	}

	if user.LastName == "" {
		return errors.New("last name is required")
	}

	if len(user.Password) < 6 {
		return errors.New("password must contain at least 6 characters")
	}

	existingUser, _ := s.Repo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("email already exists")
	}

	existingUMCN, _ := s.Repo.FindByUMCN(user.UMCN)
	if existingUMCN != nil {
		return errors.New("jmbg already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	return s.Repo.Create(user)
}

func (s *AuthService) Login(email, password string) (*models.Member, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (s *AuthService) GetMemberByID(id uint) (*dto.MemberResponse, error) {
	user, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.MemberResponse{
		ID:        user.ID,
		UMCN:      user.UMCN,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
	}, nil
}

func (s *AuthService) GetAllMembers() ([]models.Member, error) {
	return s.Repo.GetAllMembers()
}

func (s *AuthService) GetAllStudents() ([]models.Member, error) {
	return s.Repo.GetAllStudents()
}

func (s *AuthService) GetAllTeachers() ([]models.Member, error) {
	return s.Repo.GetAllTeachers()
}

func (s *AuthService) GetAllDoctors() ([]models.Member, error) {
	return s.Repo.GetAllDoctors()
}

func (s *AuthService) UpdateMember(id uint, data *dto.UpdateMemberRequest) (*dto.MemberResponse, error) {
	member, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("member not found")
	}

	if data.UMCN != "" && data.UMCN != member.UMCN {
		if !ValidateJMBG(data.UMCN) {
			return nil, errors.New("invalid jmbg")
		}

		existing, _ := s.Repo.FindByUMCN(data.UMCN)
		if existing != nil && existing.ID != id {
			return nil, errors.New("jmbg already exists")
		}

		member.UMCN = data.UMCN
	}

	if data.Email != "" && data.Email != member.Email {
		if !ValidateEmail(data.Email) {
			return nil, errors.New("invalid email")
		}

		existing, _ := s.Repo.FindByEmail(data.Email)
		if existing != nil && existing.ID != id {
			return nil, errors.New("email already exists")
		}

		member.Email = data.Email
	}

	if data.FirstName != "" {
		member.FirstName = data.FirstName
	}

	if data.LastName != "" {
		member.LastName = data.LastName
	}

	if data.Role != "" {
		member.Role = data.Role
	}

	if data.Password != "" {
		if len(data.Password) < 6 {
			return nil, errors.New("password must contain at least 6 characters")
		}

		hash, err := bcrypt.GenerateFromPassword(
			[]byte(data.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return nil, err
		}

		member.Password = string(hash)
	}

	err = s.Repo.UpdateMember(member)
	if err != nil {
		return nil, err
	}

	return &dto.MemberResponse{
		ID:        member.ID,
		UMCN:      member.UMCN,
		FirstName: member.FirstName,
		LastName:  member.LastName,
		Email:     member.Email,
		Role:      member.Role,
	}, nil
}

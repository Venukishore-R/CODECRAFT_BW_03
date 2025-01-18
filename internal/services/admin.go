package services

import (
	"fmt"

	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/models"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/repository"
)

type Admin struct {
	UserRepo    *repository.UserRepo
	UserService *UserService
}

type AdminService interface {
	GetUsers() (int, []*models.User, error)
	GetUser(email string) (int, *models.User, error)
	Create(user *models.User) (int, error)
	UpdateUser(user *models.User, email string) (int, *models.User, error)
	// Delete()
}

func NewAdminService(userRepo *repository.UserRepo) *Admin {
	return &Admin{
		UserRepo: userRepo,
	}
}

func (s *Admin) GetUsers() (int, []*models.User, error) {
	users, err := s.UserRepo.GetUsers()
	if err != nil {
		return 500, nil, err
	}

	return 200, users, nil
}

func (s *Admin) GetUser(emaail string) (int, *models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(emaail)
	if err != nil {
		return 404, nil, err
	}

	return 200, user, nil
}

func (s *Admin) Create(user *models.User) (int, error) {
	if !s.UserService.IsValidEmail(user.Email) {
		return 400, fmt.Errorf("invalid email format: %v", user.Email)
	}
	if err := s.UserRepo.CreateUser(user); err != nil {
		return 500, err
	}

	return 201, nil
}

func (s *Admin) UpdateUser(user *models.User, email string) (int, *models.User, error) {
	code, _, err := s.GetUser(email)
	if err != nil {
		return code, nil, err
	}

	result, err := s.UserRepo.Update(user, email)
	if err != nil {
		return 500, nil, err
	}

	return 200, result, nil
}

func (s *Admin) Delete(email string) (int, error) {
	code, _, err := s.GetUser(email)
	if err != nil {
		return code, err
	}

	if err := s.UserRepo.Delete(email); err != nil {
		return 500, err
	}

	return 200, nil
}

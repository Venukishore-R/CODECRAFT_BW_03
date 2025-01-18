package services

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Venukishore-R/CODECRAFT_BW_03/config"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/auth"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/models"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

type UserServiceInterface interface {
	CreateUser(user *models.User) (int, error)
	HashPassword(password string) (string, error)
	IsValidEmail(email string) bool
	Login(email, password string) (int, string, error)
	VerifyPassword(password string, hash string) (bool, error)
	UserProfile(email string) (int, *models.User, error)
}

func (s *UserService) CreateUser(user *models.User) (int, error) {
	if !s.IsValidEmail(user.Email) {
		return 400, fmt.Errorf("invalid email format")
	}

	hash, err := s.HashPassword(user.Password)
	if err != nil {
		return 500, err
	}

	user.Password = hash
	if err := s.UserRepo.CreateUser(user); err != nil {
		return 500, err
	}

	return 201, nil
}

func (s *UserService) Login(email, password string) (int, string, error) {
	if !s.IsValidEmail(email) {
		return 400, "", fmt.Errorf("invalid email format")
	}

	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return 404, "", fmt.Errorf("user not found")
	}

	isMatch, err := s.VerifyPassword(password, user.Password)
	if err != nil {
		return 500, "", err
	}

	if !isMatch {
		return 401, "", fmt.Errorf("invalid password")
	}

	config, _ := config.LoadConfig()

	auth := auth.NewAuth(
		user,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
		config.Key,
	)

	token, err := auth.GenerateToken()
	if err != nil {
		return 500, "", fmt.Errorf("unable to generate token: %v", err)
	}

	return 200, token, nil
}

func (s *UserService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("unable to generate hashed password: %v", err)
	}

	return string(hash), nil
}

func (s *UserService) VerifyPassword(password string, hash string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false, fmt.Errorf("invalid password: %v", err)
	}

	return true, nil
}

func (s *UserService) IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	reg := regexp.MustCompile(emailPattern)

	return reg.MatchString(email)
}

func (s *UserService) UserProfile(email string) (int, *models.User, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return 500, nil, err
	}

	return 200, user, nil
}

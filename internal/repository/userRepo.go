package repository

import (
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

type UserRepoInterface interface {
	CreateUser(user *models.User) error
	GetUserByEmail(id uint) (*models.User, error)
	GetUsers() ([]*models.User, error)
	Update(user *models.User, email string) (*models.User, error)
	Delete(email string) error
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.Db.Model(&models.User{}).Create(&user).Error
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	return user, r.Db.Where("email =?", email).First(&user).Error
}

func (r *UserRepo) GetUsers() ([]*models.User, error) {
	var users []*models.User
	return users, r.Db.Find(&users).Error
}

func (r *UserRepo) Update(user *models.User, email string) (*models.User, error) {
	if err := r.Db.Model(&models.User{}).Where("email =?", email).Updates(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) Delete(email string) error {
	if err := r.Db.Model(&models.User{}).Where("email =?", email).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}

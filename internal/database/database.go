package database

import (
	"fmt"

	"github.com/Venukishore-R/CODECRAFT_BW_03/config"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}

package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Key        string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	key := os.Getenv("JWT_ACCESS_TOKEN_SECRET")

	log.Println("config ", port, dbHost, dbPort, dbUser, dbPassword, dbName)
	if port == "" || dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || key == "" {
		return nil, fmt.Errorf("missing environment variables")
	}

	return &Config{
		Port:       port,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		Key:        key,
	}, nil
}

package app

import (
	"fmt"
	"log"

	"github.com/Venukishore-R/CODECRAFT_BW_03/config"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/database"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/handlers"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/repository"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/services"
	routes "github.com/Venukishore-R/CODECRAFT_BW_03/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Run() error {
	r := gin.Default()
	log.Println("s ", s.Config)
	db, err := database.ConnectDB(s.Config)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	handler := handlers.NewUserHandler(services.NewUserService(repository.NewUserRepo(db)))
	adminHandler := handlers.NewAdminHandler(services.NewAdminService(repository.NewUserRepo(db)))

	routes.Routes(r, handler, adminHandler)

	r.GET("/health", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "Ok..."}) })

	if err := r.Run(s.Config.Port); err != nil {
		return err
	}

	log.Println("Server started: ", s.Config.Port)
	return nil
}

package main

import (
	"log"

	"github.com/Venukishore-R/CODECRAFT_BW_03/config"
	"github.com/Venukishore-R/CODECRAFT_BW_03/internal/app"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
		return
	}

	server := app.NewServer(config)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

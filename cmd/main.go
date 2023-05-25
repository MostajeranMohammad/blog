package main

import (
	"log"

	"github.com/MostajeranMohammad/blog/config"
	"github.com/MostajeranMohammad/blog/internal/application"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	application.Run(cfg)
}

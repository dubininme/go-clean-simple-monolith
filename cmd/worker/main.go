package main

import (
	"log"

	"github.com/dubininme/go-clean-simple-monolith/internal/bootstrap"
	"github.com/dubininme/go-clean-simple-monolith/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	bootstrap.RunWorker(cfg)
}

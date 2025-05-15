package bootstrap

import (
	"log"

	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	"github.com/dubininme/go-clean-simple-monolith/internal/infrastructure/mysql"
)

func RunWorker(cfg *config.Config) {
	_, err := mysql.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// TODO: add worker logic here
}

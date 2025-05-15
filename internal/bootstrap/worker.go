package bootstrap

import (
	"log"
	"time"

	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	"github.com/dubininme/go-clean-simple-monolith/internal/infrastructure/mysql"
)

func RunWorker(cfg *config.Config) {
	_, err := mysql.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// TODO: add worker logic here
	go func() {
		for {
			log.Println("Worker is waiting...")
			time.Sleep(10 * time.Second)
		}
	}()

	select {}
}

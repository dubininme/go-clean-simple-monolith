package scheduler

import (
	appRepository "github.com/dubininme/go-clean-simple-monolith/internal/application/repository"
	"github.com/dubininme/go-clean-simple-monolith/internal/config"
)

type OutboxPoller struct {
	OutboxRepository appRepository.OutboxRepository
	BatchSize        int
}

func NewOutboxPoller(cfg *config.Config, outboxRepository appRepository.OutboxRepository) *OutboxPoller {
	return &OutboxPoller{
		OutboxRepository: outboxRepository,
		BatchSize:        cfg.OutboxPoller.BatchSize,
	}
}

func (p *OutboxPoller) Process() error {
	return nil
}

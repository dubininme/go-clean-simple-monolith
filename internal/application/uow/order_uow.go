package uow

import (
	"context"

	appRepository "github.com/dubininme/go-clean-simple-monolith/internal/application/repository"
	domainRepository "github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type OrderRepos struct {
	OrderRepository  domainRepository.OrderRepository
	OutboxRepository appRepository.OutboxRepository
}

type OrderUnitOfWork interface {
	DoInTx(ctx context.Context, fn func(repos *OrderRepos) error) error
}

package broker

import (
	"context"

	appRepository "github.com/dubininme/go-clean-simple-monolith/internal/application/repository"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type OrderConsumer struct {
	orderRepo repository.OrderRepository
}

func NewOrderConsumer(orderRepo repository.OrderRepository) *OrderConsumer {
	return &OrderConsumer{orderRepo: orderRepo}
}

func (c *OrderConsumer) Consume(ctx context.Context, message *appRepository.OutboxMessage) error {
	return nil
}

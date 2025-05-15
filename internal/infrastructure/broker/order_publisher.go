package broker

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type OrderPublisher struct {
	orderRepo repository.OrderRepository
}

func NewOrderPublisher(orderRepo repository.OrderRepository) *OrderPublisher {
	return &OrderPublisher{orderRepo: orderRepo}
}

func (p *OrderPublisher) Publish(ctx context.Context, order *entity.Order) error {
	return nil
}

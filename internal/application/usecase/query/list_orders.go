package query

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type ListOrdersUsecase struct {
	orderRepo repository.OrderRepository
}

func NewListOrdersUsecase(orderRepo repository.OrderRepository) *ListOrdersUsecase {
	return &ListOrdersUsecase{orderRepo: orderRepo}
}

func (uc *ListOrdersUsecase) Execute(ctx context.Context, opts repository.ListOptions) ([]entity.Order, error) {
	orders, err := uc.orderRepo.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

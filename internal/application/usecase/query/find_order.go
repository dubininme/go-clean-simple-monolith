package query

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/aggregate"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type FindOrderUsecase struct {
	orderRepo repository.OrderRepository
}

func NewFindOrderUsecase(orderRepo repository.OrderRepository) *FindOrderUsecase {
	return &FindOrderUsecase{orderRepo: orderRepo}
}

func (uc *FindOrderUsecase) Execute(ctx context.Context, id int32) (*aggregate.OrderWithItems, error) {
	extOrder, err := uc.orderRepo.FindOrderWithItemsById(ctx, id)
	if err != nil {
		return nil, err
	}

	return extOrder, nil
}

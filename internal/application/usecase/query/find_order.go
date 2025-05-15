package query

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type FindOrderUsecase struct {
	orderRepo repository.OrderRepository
}

func NewFindOrderUsecase(orderRepo repository.OrderRepository) *FindOrderUsecase {
	return &FindOrderUsecase{orderRepo: orderRepo}
}

func (uc *FindOrderUsecase) Execute(ctx context.Context, id int32) (*entity.Order, error) {
	order, err := uc.orderRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

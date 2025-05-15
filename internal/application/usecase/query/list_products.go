package query

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type ListProductsUsecase struct {
	productRepo repository.ProductRepository
}

func NewListProductsUsecase(productRepo repository.ProductRepository) *ListProductsUsecase {
	return &ListProductsUsecase{productRepo: productRepo}
}

func (uc *ListProductsUsecase) Execute(ctx context.Context, opts repository.ListOptions) ([]entity.Product, error) {
	products, err := uc.productRepo.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	return products, nil
}

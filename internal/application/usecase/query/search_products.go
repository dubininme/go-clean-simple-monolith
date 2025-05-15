package query

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type SearchProductsUsecase struct {
	searchRepo repository.ProductSearchRepository
}

func NewSearchProductsUsecase(searchRepo repository.ProductSearchRepository) *SearchProductsUsecase {
	return &SearchProductsUsecase{searchRepo: searchRepo}
}

func (uc *SearchProductsUsecase) Execute(ctx context.Context, query string) ([]entity.Product, error) {
	products, err := uc.searchRepo.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

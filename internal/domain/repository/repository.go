package repository

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/aggregate"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
)

type ListOptions struct {
	Limit     uint64
	Offset    uint64
	SortBy    string
	SortOrder string
	Filters   map[string]any
}

type ProductRepository interface {
	FindById(ctx context.Context, id int32) (*entity.Product, error)
	Save(ctx context.Context, product entity.Product) (int32, error)
	Delete(ctx context.Context, id int32) error
	List(ctx context.Context, opts ListOptions) ([]entity.Product, error)
}

type ProductSearchRepository interface {
	Search(ctx context.Context, query string) ([]entity.Product, error)
}

// TODO: split into read and write repositories
type OrderRepository interface {
	FindById(ctx context.Context, id int32) (*entity.Order, error)
	FindOrderWithItemsById(ctx context.Context, id int32) (*aggregate.OrderWithItems, error)
	List(ctx context.Context, opts ListOptions) ([]entity.Order, error)
	Save(ctx context.Context, order entity.Order) (int32, error)
	Delete(ctx context.Context, id int32) error
}

type OrderItemRepository interface {
	FindByOrderId(ctx context.Context, orderId int32) ([]entity.OrderItem, error)
	Save(ctx context.Context, item entity.OrderItem) (int32, error)
	SaveBatch(ctx context.Context, items []entity.OrderItem) error
	Delete(ctx context.Context, id int32) error
}

type ShipmentRepository interface {
	FindByOrderId(ctx context.Context, orderId int32) (*entity.Shipment, error)
	Save(ctx context.Context, shipment entity.Shipment) (int32, error)
}

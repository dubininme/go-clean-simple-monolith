package repository

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
)

type CartRepository interface {
	FindById(id string) (*entity.Cart, error)
	Save(cart *entity.Cart) error
	Delete(id string) error
	List() ([]*entity.Cart, error)
}

type ClientRepository interface {
	FindById(id string) (*entity.Client, error)
	Save(client *entity.Client) error
	Delete(id string) error
	List() ([]*entity.Client, error)
}

type CourierRepository interface {
	FindById(id string) (*entity.Courier, error)
	Save(courier *entity.Courier) error
	Delete(id string) error
	List() ([]*entity.Courier, error)
}

type OrderRepository interface {
	FindById(ctx context.Context, id string) (*entity.Order, error)
	Save(ctx context.Context, order *entity.Order) error
}

type ProductRepository interface {
	FindById(id string) (*entity.Product, error)
}

type ProductSearchRepository interface {
	Search(ctx context.Context, query string) ([]string, error)
}

type ShipmentRepository interface {
	FindById(id string) (*entity.Shipment, error)
	Save(shipment *entity.Shipment) error
	Delete(id string) error
	List() ([]*entity.Shipment, error)
}

type OrderItemRepository interface {
	FindById(id string) (*entity.OrderItem, error)
	Save(item *entity.OrderItem) error
	Delete(id string) error
	List() ([]*entity.OrderItem, error)
}

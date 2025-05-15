package shipping_provider

import (
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
)

type ShippingProviderClient struct {
}

func NewShippingProviderClient() *ShippingProviderClient {
	return &ShippingProviderClient{}
}

func (c *ShippingProviderClient) ShipOrder(order *entity.Order) error {
	// TODO: implement
	return nil
}

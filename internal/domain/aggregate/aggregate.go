package aggregate

import (
	domainEntity "github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
)

type OrderWithItems struct {
	domainEntity.Order
	Items []domainEntity.OrderItem
}

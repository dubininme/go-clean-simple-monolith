package event

import "time"

type OrderShipped struct {
	OrderId   string
	ShippedAt time.Time
}

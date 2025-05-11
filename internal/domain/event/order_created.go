package event

import "time"

type OrderCreated struct {
	OrderId   string
	ClientId  string
	CreatedAt time.Time
}

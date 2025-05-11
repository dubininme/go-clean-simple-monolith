package event

import "time"

type OrderCancelled struct {
	OrderId     string
	CancelledAt time.Time
}

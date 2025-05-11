package event

import "time"

type OrderPaid struct {
	OrderId string
	Amount  int
	PaidAt  time.Time
}

package event

import "time"

type OrderCancelled struct {
	OrderId     string
	CancelledAt time.Time
}

type OrderCreated struct {
	OrderId   string
	ClientId  string
	CreatedAt time.Time
}

type OrderPaid struct {
	OrderId string
	Amount  int
	PaidAt  time.Time
}

type OrderShipped struct {
	OrderId   string
	ShippedAt time.Time
}

type ShipmentCreated struct {
	ShipmentId string
	OrderId    string
	CreatedAt  time.Time
}

type ShipmentDelivered struct {
	ShipmentId  string
	OrderId     string
	DeliveredAt time.Time
}

type ShipmentShipped struct {
	ShipmentId string
	OrderId    string
	ShippedAt  time.Time
}

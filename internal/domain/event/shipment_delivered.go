package event

import "time"

type ShipmentDelivered struct {
	ShipmentId  string
	OrderId     string
	DeliveredAt time.Time
}

package event

import "time"

type ShipmentShipped struct {
	ShipmentId string
	OrderId    string
	ShippedAt  time.Time
}

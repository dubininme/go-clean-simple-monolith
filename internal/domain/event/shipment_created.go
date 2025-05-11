package event

import "time"

type ShipmentCreated struct {
	ShipmentId string
	OrderId    string
	CreatedAt  time.Time
}

package entity

import "time"

type Shipment struct {
	Id          string
	OrderId     string
	CourierId   string
	Status      string
	ShippedAt   *time.Time
	DeliveredAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

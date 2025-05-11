package entity

import "time"

const (
	OrderStatusPending = "pending"
	OrderStatusPaid    = "paid"
)

type Order struct {
	Id        string
	ClientId  string
	CourierId string
	Status    string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

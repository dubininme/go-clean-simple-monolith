package entity

import "time"

type OrderItem struct {
	Id        string
	OrderId   string
	ProductId string
	Quantity  int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

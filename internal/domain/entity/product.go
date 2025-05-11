package entity

import "time"

type Product struct {
	Id          string
	Name        string
	Description string
	Price       int
	SKU         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

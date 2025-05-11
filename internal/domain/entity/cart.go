package entity

import "time"

type Cart struct {
	Id        string
	ClientId  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

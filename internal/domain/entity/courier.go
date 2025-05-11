package entity

import "time"

type Courier struct {
	Id        string
	Name      string
	Phone     string
	Vehicle   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

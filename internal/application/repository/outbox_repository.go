package repository

import "time"

const (
	OutboxMessageStatusPending   = "pending"
	OutboxMessageStatusProcessed = "processed"
	OutboxMessageStatusFailed    = "failed"

	OutboxMessageTypeOrderPaid = "order_paid"
)

type OutboxMessage struct {
	Id        string
	EventType string
	Payload   []byte
	Status    string
	CreatedAt time.Time
}

type OutboxRepository interface {
	Add(msg OutboxMessage) error
}

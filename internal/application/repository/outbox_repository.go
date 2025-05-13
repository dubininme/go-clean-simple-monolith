package repository

import (
	"context"
	"time"
)

const (
	OutboxMessageStatusPending   = "pending"
	OutboxMessageStatusProcessed = "processed"
	OutboxMessageStatusFailed    = "failed"

	OutboxMessageTypeOrderPaid      = "order_paid"
	OutboxMessageTypeOrderCancelled = "order_cancelled"
)

type OutboxMessage struct {
	Id        int32
	EventType string
	Payload   []byte
	CreatedAt time.Time
}

type OutboxRepository interface {
	Add(ctx context.Context, msg OutboxMessage) error
	FetchPendingBatch(ctx context.Context, size uint64) ([]OutboxMessage, error)
	Delete(ctx context.Context, id int32) error
}

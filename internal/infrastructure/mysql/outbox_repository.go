package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	appRepository "github.com/dubininme/go-clean-simple-monolith/internal/application/repository"
)

type outboxDAO struct {
	Id        int32  `db:"id"`
	EventType string `db:"event_type"`
	Payload   []byte `db:"payload"`
	CreatedAt int64  `db:"created_at"`
}
type MysqlOutboxRepository struct {
	ext Ext
}

func NewMysqlOutboxRepository(ext Ext) appRepository.OutboxRepository {
	return &MysqlOutboxRepository{ext: ext}
}

func (r *MysqlOutboxRepository) Add(ctx context.Context, msg appRepository.OutboxMessage) error {
	queryBuilder := sq.Insert("outbox").
		Columns("event_type", "payload", "created_at").
		Values(msg.EventType, msg.Payload, msg.CreatedAt)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.ext.ExecContext(ctx, query, args...)
	return err
}

func (r *MysqlOutboxRepository) FetchPendingBatch(ctx context.Context, limit uint64) ([]appRepository.OutboxMessage, error) {
	queryBuilder := sq.Select("*").
		From("outbox").
		Where(sq.Eq{"status": appRepository.OutboxMessageStatusPending}).
		OrderBy("created_at").
		Limit(limit)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var outboxMessagesDAO []outboxDAO
	err = r.ext.SelectContext(ctx, &outboxMessagesDAO, query, args...)
	if err != nil {
		return nil, err
	}

	outboxMessages := make([]appRepository.OutboxMessage, 0, len(outboxMessagesDAO))
	for _, outboxMessageDAO := range outboxMessagesDAO {
		outboxMessages = append(outboxMessages, appRepository.OutboxMessage{
			Id:        outboxMessageDAO.Id,
			EventType: outboxMessageDAO.EventType,
			Payload:   outboxMessageDAO.Payload,
		})
	}

	return outboxMessages, nil
}

func (r *MysqlOutboxRepository) Delete(ctx context.Context, id int32) error {
	queryBuilder := sq.Delete("outbox").Where(sq.Eq{"id": id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.ext.ExecContext(ctx, query, args...)
	return err
}

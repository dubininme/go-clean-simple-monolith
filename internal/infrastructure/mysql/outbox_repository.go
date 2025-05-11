package mysql

import (
	appRepository "github.com/dubininme/go-clean-simple-monolith/internal/application/repository"
	"github.com/jmoiron/sqlx"
)

type mysqlOutboxRepository struct {
	tx *sqlx.Tx
}

func NewMysqlOutboxRepository(tx *sqlx.Tx) *mysqlOutboxRepository {
	return &mysqlOutboxRepository{tx: tx}
}

func (r *mysqlOutboxRepository) Add(msg appRepository.OutboxMessage) error {
	return nil
}

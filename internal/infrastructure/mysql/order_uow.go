package mysql

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/uow"
	"github.com/jmoiron/sqlx"
)

type mysqlOrderUow struct {
	db *sqlx.DB
}

func NewMysqlOrderUow(db *sqlx.DB) uow.OrderUnitOfWork {
	return &mysqlOrderUow{db: db}
}

func (u *mysqlOrderUow) DoInTx(ctx context.Context, fn func(repos *uow.OrderRepos) error) error {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	orderRepo := NewMysqlOrderRepository(tx)
	outboxRepo := NewMysqlOutboxRepository(tx)

	repos := &uow.OrderRepos{
		OrderRepository:  orderRepo,
		OutboxRepository: outboxRepo,
	}

	return fn(repos)
}

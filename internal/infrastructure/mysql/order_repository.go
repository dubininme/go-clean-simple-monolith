package mysql

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type mysqlOrderRepository struct {
	ext Ext
}

func NewMysqlOrderRepository(ext Ext) repository.OrderRepository {
	return &mysqlOrderRepository{ext: ext}
}

func (r *mysqlOrderRepository) FindById(ctx context.Context, id string) (*entity.Order, error) {
	var order entity.Order
	err := r.ext.GetContext(ctx, &order, "SELECT * FROM `order` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *mysqlOrderRepository) Save(ctx context.Context, order *entity.Order) error {
	_, err := r.ext.ExecContext(ctx, "INSERT INTO `order` (id, client_id, courier_id, status, amount, created_at, updated_at, deleted_at) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?) "+
		"ON DUPLICATE KEY UPDATE "+
		"client_id=VALUES(client_id), "+
		"courier_id=VALUES(courier_id), "+
		"status=VALUES(status), "+
		"amount=VALUES(amount), "+
		"updated_at=VALUES(updated_at), "+
		"deleted_at=VALUES(deleted_at)",
		order.Id, order.ClientId, order.CourierId, order.Status, order.Amount, order.CreatedAt, order.UpdatedAt, order.DeletedAt)
	return err
}

func (r *mysqlOrderRepository) Delete(id string) error {
	_, err := r.ext.ExecContext(context.Background(), "DELETE FROM `order` WHERE id = ?", id)
	return err
}

func (r *mysqlOrderRepository) List() ([]*entity.Order, error) {
	var orders []*entity.Order
	err := r.ext.SelectContext(context.Background(), &orders, "SELECT * FROM `order`")
	return orders, err
}

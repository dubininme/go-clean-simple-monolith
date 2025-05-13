package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type MysqlOrderItemRepository struct {
	ext Ext
}

func NewMysqlOrderItemRepository(ext Ext) repository.OrderItemRepository {
	return &MysqlOrderItemRepository{ext: ext}
}

type orderItemDAO struct {
	Id        int32 `db:"id"`
	OrderId   int32 `db:"order_id"`
	ProductId int32 `db:"product_id"`
	Quantity  int32 `db:"quantity"`
	Price     int64 `db:"price"`
	CreatedAt int64 `db:"created_at"`
	UpdatedAt int64 `db:"updated_at"`
}

func (r *MysqlOrderItemRepository) FindByOrderId(ctx context.Context, orderId int32) ([]entity.OrderItem, error) {
	queryBuilder := sq.Select("*").From("order_items").Where(sq.Eq{"order_id": orderId})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var daos []orderItemDAO
	err = r.ext.SelectContext(ctx, &daos, query, args...)
	if err != nil {
		return nil, err
	}

	items := make([]entity.OrderItem, 0, len(daos))
	for _, dao := range daos {
		items = append(items, toOrderItemEntity(dao))
	}
	return items, nil
}

func (r *MysqlOrderItemRepository) Save(ctx context.Context, item entity.OrderItem) (int32, error) {
	d := toOrderItemDAO(item)
	if item.Id == 0 {
		return r.insertOrderItem(ctx, d)
	} else {
		return r.updateOrderItem(ctx, d)
	}
}

// insertOrderItem builds and executes an insert query, returning the new id
func (r *MysqlOrderItemRepository) insertOrderItem(ctx context.Context, d orderItemDAO) (int32, error) {
	queryBuilder := sq.Insert("order_items").
		Columns("order_id", "product_id", "quantity", "price", "created_at", "updated_at").
		Values(d.OrderId, d.ProductId, d.Quantity, d.Price, d.CreatedAt, d.UpdatedAt)
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	result, err := r.ext.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}

// updateOrderItem builds and executes an update query, returning the id used
func (r *MysqlOrderItemRepository) updateOrderItem(ctx context.Context, d orderItemDAO) (int32, error) {
	queryBuilder := sq.Update("order_items").
		Set("order_id", d.OrderId).
		Set("product_id", d.ProductId).
		Set("quantity", d.Quantity).
		Set("price", d.Price).
		Set("updated_at", d.UpdatedAt).
		Where(sq.Eq{"id": d.Id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	_, err = r.ext.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return int32(d.Id), nil
}

// SaveBatch inserts multiple order items in a single query
func (r *MysqlOrderItemRepository) SaveBatch(ctx context.Context, items []entity.OrderItem) error {
	if len(items) == 0 {
		return nil
	}
	builder := sq.Insert("order_items").
		Columns("order_id", "product_id", "quantity", "price", "created_at", "updated_at")
	for _, item := range items {
		d := toOrderItemDAO(item)
		builder = builder.Values(d.OrderId, d.ProductId, d.Quantity, d.Price, d.CreatedAt, d.UpdatedAt)
	}
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.ext.ExecContext(ctx, query, args...)
	return err
}

func (r *MysqlOrderItemRepository) Delete(ctx context.Context, id int32) error {
	queryBuilder := sq.Delete("order_items").Where(sq.Eq{"id": id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.ext.ExecContext(ctx, query, args...)
	return err
}

func toOrderItemEntity(dao orderItemDAO) entity.OrderItem {
	return entity.OrderItem{
		Id:        dao.Id,
		OrderId:   dao.OrderId,
		ProductId: dao.ProductId,
		Quantity:  dao.Quantity,
		Price:     dao.Price,
		CreatedAt: dao.CreatedAt,
		UpdatedAt: dao.UpdatedAt,
	}
}

func toOrderItemDAO(item entity.OrderItem) orderItemDAO {
	return orderItemDAO{
		Id:        item.Id,
		OrderId:   item.OrderId,
		ProductId: item.ProductId,
		Quantity:  item.Quantity,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

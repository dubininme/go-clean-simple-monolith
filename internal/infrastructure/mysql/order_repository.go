package mysql

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/aggregate"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type MysqlOrderRepository struct {
	ext Ext
}

func NewMysqlOrderRepository(ext Ext) repository.OrderRepository {
	return &MysqlOrderRepository{ext: ext}
}

type orderDAO struct {
	Id          int32  `db:"id"`
	Status      string `db:"status"`
	Amount      int64  `db:"amount"`
	Currency    string `db:"currency"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	CancelledAt *int64 `db:"cancelled_at"`
}

func toOrderEntity(dao orderDAO) entity.Order {
	return entity.Order{
		Id:          dao.Id,
		Status:      dao.Status,
		Amount:      dao.Amount,
		Currency:    dao.Currency,
		CreatedAt:   dao.CreatedAt,
		UpdatedAt:   dao.UpdatedAt,
		CancelledAt: dao.CancelledAt,
	}
}

func toOrderDAO(order entity.Order) orderDAO {
	return orderDAO{
		Id:          order.Id,
		Status:      order.Status,
		Amount:      order.Amount,
		Currency:    order.Currency,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		CancelledAt: order.CancelledAt,
	}
}

func (r *MysqlOrderRepository) FindById(ctx context.Context, id int32) (*entity.Order, error) {
	queryBuilder := sq.Select("*").From("`order`").Where(sq.Eq{"id": id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var dao orderDAO
	err = r.ext.GetContext(ctx, &dao, query, args...)
	if err != nil {
		return nil, err
	}
	order := toOrderEntity(dao)
	return &order, nil
}

func (r *MysqlOrderRepository) FindOrderWithItemsById(ctx context.Context, id int32) (*aggregate.OrderWithItems, error) {
	type joinedRow struct {
		// order fields
		OrderId     int32  `db:"order_id"`
		Status      string `db:"status"`
		Amount      int64  `db:"amount"`
		Currency    string `db:"currency"`
		CreatedAt   int64  `db:"created_at"`
		UpdatedAt   int64  `db:"updated_at"`
		CancelledAt *int64 `db:"cancelled_at"`
		// item fields
		ItemId        *int32 `db:"item_id"`
		ProductId     *int32 `db:"product_id"`
		Quantity      *int32 `db:"quantity"`
		Price         *int64 `db:"price"`
		ItemCreatedAt *int64 `db:"item_created_at"`
		ItemUpdatedAt *int64 `db:"item_updated_at"`
	}

	query, args, err := sq.
		Select(
			"`order`.id AS order_id",
			"`order`.status",
			"`order`.amount",
			"`order`.currency",
			"`order`.created_at",
			"`order`.updated_at",
			"`order`.cancelled_at",
			"`order_item`.id AS item_id",
			"`order_item`.product_id",
			"`order_item`.quantity",
			"`order_item`.price",
			"`order_item`.created_at AS item_created_at",
			"`order_item`.updated_at AS item_updated_at",
		).
		From("`order`").
		LeftJoin("`order_item` ON `order_item`.`order_id` = `order`.`id`").
		Where(sq.Eq{"`order`.id": id}).
		PlaceholderFormat(sq.Question).
		ToSql()
	if err != nil {
		return nil, err
	}

	var rows []joinedRow
	err = r.ext.SelectContext(ctx, &rows, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || len(rows) == 0 {
			return nil, nil
		}
		return nil, err
	}
	order := entity.Order{
		Id:          rows[0].OrderId,
		Status:      rows[0].Status,
		Amount:      rows[0].Amount,
		Currency:    rows[0].Currency,
		CreatedAt:   rows[0].CreatedAt,
		UpdatedAt:   rows[0].UpdatedAt,
		CancelledAt: rows[0].CancelledAt,
	}
	items := make([]entity.OrderItem, 0, len(rows))
	for _, row := range rows {
		if row.ItemId == nil {
			continue // skip if order has no items
		}

		items = append(items, entity.OrderItem{
			Id:        *row.ItemId,
			OrderId:   rows[0].OrderId,
			ProductId: *row.ProductId,
			Quantity:  *row.Quantity,
			Price:     *row.Price,
			CreatedAt: *row.ItemCreatedAt,
			UpdatedAt: *row.ItemUpdatedAt,
		})
	}

	return &aggregate.OrderWithItems{
		Order: order,
		Items: items,
	}, nil
}

func (r *MysqlOrderRepository) Save(ctx context.Context, order entity.Order) (int32, error) {
	d := toOrderDAO(order)
	if d.Id == 0 {
		return r.insertOrder(ctx, d)
	} else {
		return r.upsertOrder(ctx, d)
	}
}

// insertOrder builds and executes an insert query, returning the new id
func (r *MysqlOrderRepository) insertOrder(ctx context.Context, d orderDAO) (int32, error) {
	queryBuilder := sq.Insert("`order`").
		Columns("status", "amount", "currency", "created_at", "updated_at", "cancelled_at").
		Values(d.Status, d.Amount, d.Currency, d.CreatedAt, d.UpdatedAt, d.CancelledAt)

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

// upsertOrder builds and executes an upsert query, returning the id used
func (r *MysqlOrderRepository) upsertOrder(ctx context.Context, d orderDAO) (int32, error) {
	queryBuilder := sq.Insert("`order`").
		Columns("id", "status", "amount", "currency", "created_at", "updated_at", "cancelled_at").
		Values(d.Id, d.Status, d.Amount, d.Currency, d.CreatedAt, d.UpdatedAt, d.CancelledAt).
		Suffix(`ON DUPLICATE KEY UPDATE
			status=VALUES(status),
			amount=VALUES(amount),
			currency=VALUES(currency),
			created_at=VALUES(created_at),
			updated_at=VALUES(updated_at),
			cancelled_at=VALUES(cancelled_at)`)
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

func (r *MysqlOrderRepository) Delete(ctx context.Context, id int32) error {
	queryBuilder := sq.Delete("`order`").Where(sq.Eq{"id": id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.ext.ExecContext(ctx, query, args...)
	return err
}

func (r *MysqlOrderRepository) List(ctx context.Context, opts repository.ListOptions) ([]entity.Order, error) {
	queryBuilder := sq.Select("*").From("`order`")

	if opts.Limit > 0 {
		queryBuilder = queryBuilder.Limit(opts.Limit)
	}
	if opts.Offset > 0 {
		queryBuilder = queryBuilder.Offset(opts.Offset)
	}

	if opts.SortBy != "" {
		orderBy := opts.SortBy
		if opts.SortOrder == "ASC" || opts.SortOrder == "DESC" {
			orderBy += " " + opts.SortOrder
		}
		queryBuilder = queryBuilder.OrderBy(orderBy)
	}

	if len(opts.Filters) > 0 {
		for k, v := range opts.Filters {
			queryBuilder = queryBuilder.Where(sq.Eq{k: v})
		}
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var daos []orderDAO
	err = r.ext.GetContext(ctx, &daos, query, args...)
	if err != nil {
		return nil, err
	}
	orders := make([]entity.Order, 0, len(daos))
	for _, dao := range daos {
		order := toOrderEntity(dao)
		orders = append(orders, order)
	}
	return orders, nil
}

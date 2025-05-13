package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type MysqlShipmentRepository struct {
	ext Ext
}

func NewMysqlShipmentRepository(ext Ext) repository.ShipmentRepository {
	return &MysqlShipmentRepository{ext: ext}
}

type shipmentDAO struct {
	Id        int32  `db:"id"`
	OrderId   int32  `db:"order_id"`
	Address   string `db:"address"`
	Phone     string `db:"phone"`
	Status    string `db:"status"`
	ShippedAt *int64 `db:"shipped_at"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func (r *MysqlShipmentRepository) FindByOrderId(ctx context.Context, orderId int32) (*entity.Shipment, error) {
	queryBuilder := sq.Select("*").From("shipments").Where(sq.Eq{"order_id": orderId})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var dao shipmentDAO
	err = r.ext.GetContext(ctx, &dao, query, args...)
	if err != nil {
		return nil, err
	}

	shipment := toShipmentEntity(dao)
	return &shipment, nil
}

func (r *MysqlShipmentRepository) Save(ctx context.Context, shipment entity.Shipment) (int32, error) {
	d := toShipmentDAO(shipment)
	if shipment.Id == 0 {
		return r.insertShipment(ctx, d)
	} else {
		return r.updateShipment(ctx, d)
	}
}

// insertShipment builds and executes an insert query, returning the new id
func (r *MysqlShipmentRepository) insertShipment(ctx context.Context, d shipmentDAO) (int32, error) {
	queryBuilder := sq.Insert("shipments").
		Columns("order_id", "address", "phone", "status", "shipped_at", "created_at", "updated_at").
		Values(d.OrderId, d.Address, d.Phone, d.Status, d.ShippedAt, d.CreatedAt, d.UpdatedAt)
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

// updateShipment builds and executes an update query, returning the id used
func (r *MysqlShipmentRepository) updateShipment(ctx context.Context, d shipmentDAO) (int32, error) {
	queryBuilder := sq.Update("shipments").
		Set("order_id", d.OrderId).
		Set("address", d.Address).
		Set("phone", d.Phone).
		Set("status", d.Status).
		Set("shipped_at", d.ShippedAt).
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

func toShipmentEntity(dao shipmentDAO) entity.Shipment {
	return entity.Shipment{
		Id:        dao.Id,
		OrderId:   dao.OrderId,
		Address:   dao.Address,
		Phone:     dao.Phone,
		Status:    dao.Status,
		ShippedAt: dao.ShippedAt,
		CreatedAt: dao.CreatedAt,
		UpdatedAt: dao.UpdatedAt,
	}
}

func toShipmentDAO(shipment entity.Shipment) shipmentDAO {
	return shipmentDAO{
		Id:        shipment.Id,
		OrderId:   shipment.OrderId,
		Address:   shipment.Address,
		Phone:     shipment.Phone,
		Status:    shipment.Status,
		ShippedAt: shipment.ShippedAt,
		CreatedAt: shipment.CreatedAt,
		UpdatedAt: shipment.UpdatedAt,
	}
}

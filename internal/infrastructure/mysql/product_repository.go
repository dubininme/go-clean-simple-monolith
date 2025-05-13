package mysql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
)

type MysqlProductRepository struct {
	ext Ext
}

func NewMysqlProductRepository(ext Ext) repository.ProductRepository {
	return &MysqlProductRepository{ext: ext}
}

type productDAO struct {
	Id          int32  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
	SKU         string `db:"sku"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
}

func (r *MysqlProductRepository) FindById(ctx context.Context, id int32) (*entity.Product, error) {
	queryBuilder := sq.Select("*").From("product").Where(sq.Eq{"id": id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var dao productDAO
	err = r.ext.GetContext(ctx, &dao, query, args...)
	if err != nil {
		return nil, err
	}
	product := toProductEntity(dao)
	return &product, nil
}

func (r *MysqlProductRepository) Save(ctx context.Context, product entity.Product) (int32, error) {
	d := toProductDAO(product)
	if product.Id == 0 {
		return r.insertProduct(ctx, d)
	} else {
		return r.updateProduct(ctx, d)
	}
}

// insertProduct builds and executes an insert query, returning the new id
func (r *MysqlProductRepository) insertProduct(ctx context.Context, d productDAO) (int32, error) {
	queryBuilder := sq.Insert("product").
		Columns("name", "description", "price", "sku", "created_at", "updated_at").
		Values(d.Name, d.Description, d.Price, d.SKU, d.CreatedAt, d.UpdatedAt)
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

// updateProduct builds and executes an update query, returning the id used
func (r *MysqlProductRepository) updateProduct(ctx context.Context, d productDAO) (int32, error) {
	queryBuilder := sq.Update("product").
		Set("name", d.Name).
		Set("description", d.Description).
		Set("price", d.Price).
		Set("sku", d.SKU).
		Set("created_at", d.CreatedAt).
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

func (r *MysqlProductRepository) Delete(ctx context.Context, id int32) error {
	queryBuilder := sq.Delete("product").Where(sq.Eq{"id": id})
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}
	_, err = r.ext.ExecContext(ctx, query, args...)
	return err
}

func (r *MysqlProductRepository) List(ctx context.Context, opts repository.ListOptions) ([]entity.Product, error) {
	queryBuilder := sq.Select("*").From("product")

	if opts.Limit > 0 {
		queryBuilder = queryBuilder.Limit(opts.Limit)
	}
	if opts.Offset > 0 {
		queryBuilder = queryBuilder.Offset(opts.Offset)
	}

	if opts.SortBy != "" {
		queryBuilder = queryBuilder.OrderBy(opts.SortBy)
	}
	if opts.SortOrder != "" {
		queryBuilder = queryBuilder.OrderBy(opts.SortOrder)
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

	var daos []productDAO
	err = r.ext.SelectContext(ctx, &daos, query, args...)
	if err != nil {
		return nil, err
	}
	products := make([]entity.Product, len(daos))
	for i, dao := range daos {
		products[i] = toProductEntity(dao)
	}
	return products, nil
}

func toProductEntity(dao productDAO) entity.Product {
	return entity.Product{
		Id:          dao.Id,
		Name:        dao.Name,
		Description: dao.Description,
		Price:       dao.Price,
		SKU:         dao.SKU,
		CreatedAt:   dao.CreatedAt,
		UpdatedAt:   dao.UpdatedAt,
	}
}

func toProductDAO(product entity.Product) productDAO {
	return productDAO{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		SKU:         product.SKU,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

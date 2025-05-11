package mysql

import (
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
	"github.com/jmoiron/sqlx"
)

type mysqlClientRepository struct {
	db *sqlx.DB
}

func NewMysqlClientRepository(db *sqlx.DB) repository.ClientRepository {
	return &mysqlClientRepository{db: db}
}

func (r *mysqlClientRepository) FindById(id string) (*entity.Client, error) {
	var client entity.Client
	err := r.db.Get(&client, "SELECT * FROM client WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *mysqlClientRepository) Save(client *entity.Client) error {
	_, err := r.db.Exec(`INSERT INTO client (id, name, email, phone, created_at, updated_at, deleted_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			name=VALUES(name),
			email=VALUES(email),
			phone=VALUES(phone),
			updated_at=VALUES(updated_at),
			deleted_at=VALUES(deleted_at)
	`, client.Id, client.Name, client.Email, client.Phone, client.CreatedAt, client.UpdatedAt, client.DeletedAt)
	return err
}

func (r *mysqlClientRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM client WHERE id = ?", id)
	return err
}

func (r *mysqlClientRepository) List() ([]*entity.Client, error) {
	var clients []*entity.Client
	err := r.db.Select(&clients, "SELECT * FROM client")
	return clients, err
}

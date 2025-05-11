package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Ext interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
}

func NewDB(cfg config.DBConfig) (*sqlx.DB, error) {
	dsn := buildDSN(cfg)
	log.Printf("Connecting to: %s:%s@tcp(%s:%d)/%s", cfg.User, "******", cfg.Host, cfg.Port, cfg.Name)

	db, err := sqlx.ConnectContext(context.Background(), "mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.ConnectionsCount)
	db.SetMaxIdleConns(cfg.ConnectionsCount)
	return db, nil
}

func buildDSN(cfg config.DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
}

package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgres(cfg PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.GetConnectionString())
	if err != nil {
		return nil, errors.Wrap(err, "failed to load the database")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping to the database")
	}

	db.SetMaxOpenConns(cfg.GetMaxPoolSize())
	if cfg.GetMaxIdleConnections() != 0 {
		db.SetMaxIdleConns(cfg.GetMaxIdleConnections())
	}
	return db, nil
}

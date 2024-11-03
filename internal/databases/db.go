package databases

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/stdlib"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"github.com/zuekveer/BWA/internal/config"
	"github.com/zuekveer/BWA/internal/logger"
)

type (
	closeFn func() error
	DB      = sql.DB
)

func NewDB(cfg *config.DB, log logger.Logger) (*sql.DB, closeFn, error) {
	connCfg, err := pgx.ParseURI(cfg.URI)
	if err != nil {
		return nil, nil, fmt.Errorf("parse URI: %w", err)
	}

	db := stdlib.OpenDB(connCfg)
	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to ping to the databases: %w", err)
	}
	log.Info("Successfully connected to the databases")
	return db, db.Close, nil
}

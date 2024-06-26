package storage

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"banner_service/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func connection(databaseDSN string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		return nil, fmt.Errorf("cannot open DataBase: %w", err)
	}

	return db, nil
}

func newStorage(conf *config.Storage) (*storageData, error) {
	db, err := connection(conf.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("cannot connection database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return &storageData{db: db}, nil
}

func Init(conf *config.Storage) (Storage, error) {
	st, err := newStorage(conf)
	if err != nil {
		return nil, fmt.Errorf("cannot create data base: %w", err)
	}

	_, err = goose.GetDBVersion(st.db)
	if err != nil {
		return nil, err
	}

	err = goose.Up(st.db, "./internal/migration")
	if err != nil && err != goose.ErrNoNextVersion {
		return nil, err
	}

	return st, nil
}

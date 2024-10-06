package db

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

const MigrationPath = "file://migrations"

// Создание подключения к БД.
func NewDB(dsn string) (*DB, error) {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	slog.Info("Connected to db pool")
	err = MigrateUp(dsn)
	if err != nil {
		return nil, err
	}
	db := &DB{dbpool}
	return db, nil
}

// Миграция таблиц БД.
func MigrateUp(dsn string) error {
	m, err := migrate.New(MigrationPath, dsn)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}
	slog.Info("Migrated tables successfully")
	return nil
}

// Закрытие соединения.
func CloseConnection(dbpool *DB) {
	dbpool.Close()
}

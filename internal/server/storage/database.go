package storage

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/pisarevaa/gophkeeper/internal/server/model"

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

func (dbpool *DB) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := dbpool.QueryRow(ctx, "SELECT id, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (dbpool *DB) GetUserByID(ctx context.Context, userID int64) (model.User, error) {
	var user model.User
	err := dbpool.QueryRow(ctx, "SELECT id, email, password, created_at FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (dbpool *DB) RegisterUser(ctx context.Context, email string, passwordHash string) (model.User, error) {
	var user model.User
	err := dbpool.QueryRow(ctx, "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING (id, email, password, created_at)", email, passwordHash).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Закрытие соединения.
func (dbpool *DB) CloseConnection() {
	dbpool.Close()
}

package db

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
)

// Получение пользователя по почте.
func (dbpool *DB) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := dbpool.QueryRow(ctx, "SELECT id, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Получение пользователя по ID.
func (dbpool *DB) GetUserByID(ctx context.Context, userID int64) (model.User, error) {
	var user model.User
	err := dbpool.QueryRow(ctx, "SELECT id, email, password, created_at FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Регистарция нового пользователя.
func (dbpool *DB) RegisterUser(ctx context.Context, email string, passwordHash string) (model.User, error) {
	var user model.User
	err := dbpool.QueryRow(ctx, "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password, created_at", email, passwordHash).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

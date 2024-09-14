package storage

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

// Получение всех данных пользователя по ID.
func (dbpool *DB) GetDataByUserID(ctx context.Context, userID int64) ([]model.Keeper, error) {
	var d []model.Keeper
	rows, err := dbpool.Query(
		ctx,
		"SELECT id, name, data, type, user_id, created_at, updated_at FROM keeper WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return d, err
	}
	defer rows.Close()
	for rows.Next() {
		var o model.Keeper
		err = rows.Scan(&o.ID, &o.Name, &o.Data, &o.Type, &o.UserID, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return d, err
		}
		d = append(d, o)
	}
	return d, nil
}

// Получение данных по ID.
func (dbpool *DB) GetDataByID(ctx context.Context, dataID int64) (model.Keeper, error) {
	var d model.Keeper
	err := dbpool.QueryRow(
		ctx,
		"SELECT id, name, data, type, user_id, created_at, updated_at FROM keeper WHERE id = $1",
		dataID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Добавление данных.
func (dbpool *DB) AddData(ctx context.Context, keeper model.AddKeeper, userID int64) (model.Keeper, error) {
	var d model.Keeper
	err := dbpool.QueryRow(
		ctx,
		"INSERT INTO keeper (name, data, type, user_id) VALUES ($1, $2, $3, $4) RETURNING id, name, data, type, user_id, created_at, updated_at",
		keeper.Name,
		keeper.Data,
		keeper.Type,
		userID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Обновление данных по ID.
func (dbpool *DB) UpdateData(ctx context.Context, keeper model.AddKeeper, dataID int64) (model.Keeper, error) {
	var d model.Keeper
	err := dbpool.QueryRow(
		ctx,
		"UPDATE keeper SET name = $1, data = $2, type = $3 WHERE id = $4 RETURNING id, name, data, type, user_id, created_at, updated_at",
		keeper.Name,
		keeper.Data,
		keeper.Type,
		dataID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Удаление данных по ID.
func (dbpool *DB) DeleteData(ctx context.Context, dataID int64) (model.Keeper, error) {
	var d model.Keeper
	err := dbpool.QueryRow(
		ctx,
		"DELETE keeper WHERE id = $1 RETURNING id, name, data, type, user_id, created_at, updated_at",
		dataID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	return d, nil
}

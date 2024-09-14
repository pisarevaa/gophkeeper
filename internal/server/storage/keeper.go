package storage

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/pisarevaa/gophkeeper/internal/server/model"
)

// Получение всех данных пользователя по ID.
func (dbpool *DB) GetDataByUserId(ctx context.Context, userID int64) ([]model.Keeper, error) {
	var d []model.Keeper
	rows, err := dbpool.Query(
		ctx,
		"SELECT id, data, type, user_id, created_at, updated_at FROM data_text WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return d, err
	}
	defer rows.Close()
	for rows.Next() {
		var o model.Keeper
		err = rows.Scan(&o.ID, &o.Data, &o.Type, &o.UserID, &o.CreatedAt, &o.UpdatedAt)
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
		"SELECT id, data, type, user_id, created_at, updated_at FROM data_text WHERE id = $1",
		dataID,
	).Scan(&d.ID, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Добавление новых данных.
func (dbpool *DB) AddData(ctx context.Context, keeper model.AddKeeper, userID int64) (model.Keeper, error) {
	var d model.Keeper
	err := dbpool.QueryRow(
		ctx,
		"INSERT INTO data_text (data, type, user_id) VALUES ($1, $2) RETURNING id, data, type, user_id, created_at, updated_at",
		keeper.Data,
		keeper.Type,
		userID,
	).Scan(&d.ID, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
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
		"UPDATE data_text SET data = $1 WHERE id = $2 RETURNING id, data, type, user_id, created_at, updated_at",
		keeper.Data, keeper.Type, dataID,
	).Scan(&d.ID, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
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
		"DELETE data_text WHERE id = $1 RETURNING id, data, type, user_id, created_at, updated_at",
		dataID,
	).Scan(&d.ID, &d.Data, &d.Type, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	return d, nil
}

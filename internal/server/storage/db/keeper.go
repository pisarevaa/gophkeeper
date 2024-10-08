package db

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/pisarevaa/gophkeeper/internal/shared/model"
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
		var dataType string
		err = rows.Scan(&o.ID, &o.Name, &o.Data, &dataType, &o.UserID, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return d, err
		}
		err = o.Type.SetValue(dataType)
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
	var dataType string
	err := dbpool.QueryRow(
		ctx,
		"SELECT id, name, data, object_id, filename, type, user_id, created_at, updated_at FROM keeper WHERE id = $1",
		dataID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.ObjectID, &d.FileName, &dataType, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	err = d.Type.SetValue(dataType)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Добавление данных.
func (dbpool *DB) AddData(ctx context.Context, keeper model.AddKeeper, userID int64) (model.Keeper, error) {
	var d model.Keeper
	var dataType string
	err := dbpool.QueryRow(
		ctx,
		"INSERT INTO keeper (name, data, object_id, filename, type, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, data, object_id, filename, type, user_id, created_at, updated_at",
		keeper.Name,
		keeper.Data,
		keeper.ObjectID,
		keeper.FileName,
		keeper.Type.String(),
		userID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.ObjectID, &d.FileName, &dataType, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	err = d.Type.SetValue(dataType)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Обновление данных по ID.
func (dbpool *DB) UpdateData(ctx context.Context, keeper model.AddKeeper, dataID int64) (model.Keeper, error) {
	var d model.Keeper
	var dataType string
	err := dbpool.QueryRow(
		ctx,
		"UPDATE keeper SET name = $1, data = $2, object_id = $3, filename = $4, type = $5 WHERE id = $6 RETURNING id, name, data, object_id, filename, type, user_id, created_at, updated_at",
		keeper.Name,
		keeper.Data,
		keeper.ObjectID,
		keeper.FileName,
		keeper.Type.String(),
		dataID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.ObjectID, &d.FileName, &dataType, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	err = d.Type.SetValue(dataType)
	if err != nil {
		return d, err
	}
	return d, nil
}

// Удаление данных по ID.
func (dbpool *DB) DeleteData(ctx context.Context, dataID int64) (model.Keeper, error) {
	var d model.Keeper
	var dataType string
	err := dbpool.QueryRow(
		ctx,
		"DELETE FROM keeper WHERE id = $1 RETURNING id, name, data, object_id, filename, type, user_id, created_at, updated_at",
		dataID,
	).Scan(&d.ID, &d.Name, &d.Data, &d.ObjectID, &d.FileName, &dataType, &d.UserID, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return d, err
	}
	err = d.Type.SetValue(dataType)
	if err != nil {
		return d, err
	}
	return d, nil
}

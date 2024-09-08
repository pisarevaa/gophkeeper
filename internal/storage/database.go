package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type DB struct {
	*pgxpool.Pool
}

func NewDB(dsn string, logger *zap.SugaredLogger) *DB {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Error("Unable to create connection db pool: %v", err)
		return nil
	}
	logger.Info("Connected to db pool")
	db := &DB{dbpool}
	return db
}

// Сохранение метрики.
func (dbpool *DB) StoreMetric(ctx context.Context, metric Metrics) error {
	now := time.Now()
	_, err := dbpool.Exec(ctx, `
			INSERT INTO metrics (id, type, delta, value, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (id) DO UPDATE
			SET
			type = excluded.type,
			delta = (SELECT delta FROM metrics WHERE id = $6) + excluded.delta,
			value = excluded.value,
			updated_at = excluded.updated_at
		`, metric.ID, metric.MType, metric.Delta, metric.Value, now, metric.ID)
	if err != nil {
		return err
	}
	return nil
}

// Сохранение метрик.
func (dbpool *DB) StoreMetrics(ctx context.Context, metrics []Metrics) error {
	tx, errTx := dbpool.Begin(ctx)
	if errTx != nil {
		return errTx
	}
	for _, metric := range metrics {
		err := dbpool.StoreMetric(
			ctx,
			metric,
		)
		if err != nil {
			return err
		}
	}
	errTx = tx.Commit(ctx)
	if errTx != nil {
		return errTx
	}
	return nil
}

// Получение метрики.
func (dbpool *DB) GetMetric(ctx context.Context, id string, mtype string) (Metrics, error) {
	var m Metrics
	err := dbpool.QueryRow(ctx, "SELECT id, type, delta, value FROM metrics WHERE id = $1 and type = $2", id, mtype).
		Scan(&m.ID, &m.MType, &m.Delta, &m.Value)
	if err != nil {
		return m, err
	}
	return m, nil
}

// Получение метрик.
func (dbpool *DB) GetAllMetrics(ctx context.Context) ([]Metrics, error) {
	var metrics []Metrics
	rows, err := dbpool.Query(ctx, "SELECT id, type, delta, value FROM metrics")
	if err != nil {
		return []Metrics{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var m Metrics
		err = rows.Scan(&m.ID, &m.MType, &m.Delta, &m.Value)
		if err != nil {
			return []Metrics{}, err
		}
		metrics = append(metrics, m)
	}
	return metrics, nil
}

// Пинг хранилища.
func (dbpool *DB) Ping(ctx context.Context) error {
	var one int
	err := dbpool.QueryRow(ctx, "select 1").Scan(&one)
	if err != nil {
		return err
	}
	return nil
}

func (dbpool *DB) CloseConnection() {
	dbpool.Close()
}

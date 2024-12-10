package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPGXPool(connString string) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse db config: %w", err)
	}

	dbConfig.MaxConns = 10
	dbConfig.MinConns = 2
	healthCheckPeriod := 30 * time.Second
	dbConfig.HealthCheckPeriod = healthCheckPeriod

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create db pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()

		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return pool, nil
}

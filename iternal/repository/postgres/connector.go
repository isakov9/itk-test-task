package postgres

import (
	"context"
	"fmt"
	"itk-test-task/iternal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConnector struct {
	Pool *pgxpool.Pool
}

func NewDBConnector(cfg *config.Config) (*DBConnector, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase)

	serverConfig, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, fmt.Errorf("failed to parse dsn: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), serverConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}

	return &DBConnector{Pool: pool}, nil
}

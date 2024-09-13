package storage

import (
	"context"
	"fmt"
	"github.com/DanilMargaryan/microservices/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.PostgreSQL) (*Storage, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	parseConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	parseConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe
	pool, err := pgxpool.NewWithConfig(ctx, parseConfig)
	if err != nil {
		return nil, err
	}

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}
	defer conn.Release()

	fmt.Println("Connected!")
	return &Storage{pool}, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}

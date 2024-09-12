package storage

import (
	"context"
	"fmt"
	"github.com/DanilMargaryan/microservices/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.PostgreSQL) (*Storage, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected!")
	return &Storage{pool}, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}

package storage

import (
	"database/sql"
	"fmt"
	"github.com/DanilMargaryan/microservices/internal/config"
	_ "github.com/lib/pq"
)

// Структура, содержащая подключение к базе данных
type Storage struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Storage, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected!")
	return &Storage{DB: db}, nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}

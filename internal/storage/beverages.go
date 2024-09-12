package storage

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

type Beverage struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

const (
	getAllBeverages = "SELECT name, type, price FROM beverages"
	getBeverage     = "SELECT name, type, price FROM beverages WHERE id=$1"
	createBeverage  = "INSERT INTO beverages (name, type, price, description) VALUES($1, $2, $3, $4)"
)

// GetAllBeverages godoc
// @Summary Получить все напитки
// @Description Возвращает список всех напитков из базы данных
// @Tags beverages
// @Produce  json
// @Success 200 {array} storage.Beverage
// @Failure 500 {string} string "Ошибка при получении данных"
// @Router /beverages [get]
func (s *Storage) GetAllBeverages(ctx context.Context) ([]Beverage, error) {
	const op = "storage.beverage.GetAllBeverages"

	rows, err := s.pool.Query(ctx, getAllBeverages)
	if err != nil {
		return nil, errors.Errorf("%s: %v", op, err)
	}
	defer rows.Close()

	var beverages []Beverage
	for rows.Next() {
		var beverage Beverage
		err := rows.Scan(&beverage.Name, &beverage.Type, &beverage.Price)
		if err != nil {
			return nil, errors.Errorf("%s: %v", op, err)
		}
		beverages = append(beverages, beverage)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Errorf("%s: %v", op, err)
	}

	return beverages, nil
}

// GetBeverage godoc
// @Summary Получить напиток по id
// @Description Возвращает данные напитка по его идентификатору
// @Tags beverages
// @Produce  json
// @Param id path int true "ID напитка"
// @Success 200 {object} storage.Beverage
// @Failure 404 {string} string "Напиток не найден"
// @Failure 500 {string} string "Ошибка при получении данных"
// @Router /beverage/{id} [get]
func (s *Storage) GetBeverage(ctx context.Context, id int) (*Beverage, error) {
	const op = "storage.beverage.GetBeverage"

	row := s.pool.QueryRow(ctx, getBeverage, id)
	var beverage Beverage
	err := row.Scan(&beverage.Name, &beverage.Type, &beverage.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Errorf("%s: %v", op, err)
		}
		return nil, errors.Errorf("%s: %v", op, err)
	}

	return &beverage, nil
}

// CreateBeverage godoc
// @Summary Добавить новый напиток
// @Description Добавляет новый напиток в базу данных
// @Tags beverages
// @Accept  json
// @Produce  json
// @Param beverage body storage.Beverage true "Данные напитка"
// @Success 201 {string} string "Напиток успешно добавлен!"
// @Failure 400 {string} string "Ошибка при парсинге тела запроса"
// @Failure 500 {string} string "Ошибка при добавлении напитка в базу данных"
// @Router /beverage [post]
func (s *Storage) CreateBeverage(ctx context.Context, newBeverage Beverage) error {
	const op = "storage.beverage.CreateBeverage"

	_, err := s.pool.Exec(ctx, createBeverage, newBeverage.Name, newBeverage.Type, newBeverage.Price, newBeverage.Description)
	if err != nil {
		// Возвращаем ошибку с пояснением
		return errors.Errorf("%s: %v", op, err)
	}

	return nil
}

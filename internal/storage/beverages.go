package storage

import (
	"context"
	"database/sql"
	"log"
)

// Структура для напитков
type Beverage struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

// GetAllBeverages godoc
// @Summary Получить все напитки
// @Description Возвращает список всех напитков из базы данных
// @Tags beverages
// @Produce  json
// @Success 200 {array} storage.Beverage
// @Failure 500 {string} string "Ошибка при получении данных"
// @Router /beverages [get]
func (s *Storage) GetAllBeverages(ctx context.Context) ([]Beverage, error) {
	rows, err := s.DB.QueryContext(ctx, "SELECT name, type, price FROM beverages")
	if err != nil {
		log.Println("Ошибка при выполнении запроса к базе данных:", err)
		return nil, err
	}
	defer rows.Close()

	var beverages []Beverage
	for rows.Next() {
		var beverage Beverage
		err := rows.Scan(&beverage.Name, &beverage.Type, &beverage.Price)
		if err != nil {
			log.Println("Ошибка при чтении данных:", err)
			return nil, err
		}
		beverages = append(beverages, beverage)
	}

	// Проверка ошибок после цикла
	if err = rows.Err(); err != nil {
		log.Println("Ошибка при итерации по строкам:", err)
		return nil, err
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
	row := s.DB.QueryRowContext(ctx, "SELECT name, type, price FROM beverages WHERE id=$1", id)
	var beverage Beverage
	err := row.Scan(&beverage.Name, &beverage.Type, &beverage.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Напиток с таким id не найден")
			return nil, err
		}
		log.Println("Ошибка при чтении данных:", err)
		return nil, err
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
	stmt, err := s.DB.PrepareContext(ctx, "INSERT INTO beverages (name, type, price, description) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Println("Ошибка при подготовке запроса:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, newBeverage.Name, newBeverage.Type, newBeverage.Price, newBeverage.Description)
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return err
	}

	return nil
}

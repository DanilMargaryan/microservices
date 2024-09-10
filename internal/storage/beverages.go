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

// Функция для получения всех напитков из таблицы
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

// Функция для получения напитка
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

// Функция для добавления напитка
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

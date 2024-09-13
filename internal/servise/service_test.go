package servise

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DanilMargaryan/microservices/internal/storage"
	"github.com/DanilMargaryan/microservices/internal/storage/mocks"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllBeverages_Success(t *testing.T) {
	// Создаем мок для StorageInterface
	mockStorage := new(mocks.StorageInterface)
	handler := NewHandler(mockStorage)

	// Подготовка данных напитков
	beverages := []storage.Beverage{
		{Name: "Cola", Type: "Soft Drink", Price: 1.99},
		{Name: "Pepsi", Type: "Soft Drink", Price: 1.89},
	}

	// Настраиваем мок: ожидаем, что GetAllBeverages вернет список напитков
	mockStorage.On("GetAllBeverages", mock.Anything).Return(beverages, nil)

	// Создаем Fiber приложение
	app := fiber.New()

	// Регистрируем хендлер
	app.Get("/beverages", handler.GetAllBeverages)

	// Выполняем HTTP-запрос
	req := httptest.NewRequest(http.MethodGet, "/beverages", nil)
	resp, _ := app.Test(req)

	// Проверяем успешный статус
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockStorage.AssertExpectations(t)
}

func TestGetAllBeverages_Failure(t *testing.T) {
	// Создаем мок для StorageInterface
	mockStorage := new(mocks.StorageInterface)
	handler := NewHandler(mockStorage)

	// Настраиваем мок: ожидаем, что GetAllBeverages вернет ошибку
	mockStorage.On("GetAllBeverages", mock.Anything).Return(nil, assert.AnError)

	// Создаем Fiber приложение
	app := fiber.New()

	// Регистрируем хендлер
	app.Get("/beverages", handler.GetAllBeverages)

	// Выполняем HTTP-запрос
	req := httptest.NewRequest(http.MethodGet, "/beverages", nil)
	resp, _ := app.Test(req)

	// Проверяем, что был возвращен код 500 (Internal Server Error)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockStorage.AssertExpectations(t)
}

func TestCreateBeverage_Success(t *testing.T) {
	// Создаем мок для StorageInterface
	mockStorage := new(mocks.StorageInterface)
	handler := NewHandler(mockStorage)

	// Подготовка данных напитка
	newBeverage := storage.Beverage{
		Name:        "Cola",
		Type:        "Soft Drink",
		Price:       1.99,
		Description: "A refreshing drink",
	}

	// Настраиваем мок: ожидаем, что CreateBeverage будет вызван с newBeverage и не вернет ошибку
	mockStorage.On("CreateBeverage", mock.Anything, newBeverage).Return(nil)

	// Создаем Fiber приложение
	app := fiber.New()

	// Регистрируем хендлер
	app.Post("/beverage", handler.CreateBeverage)

	// Подготавливаем HTTP-запрос
	reqBody, _ := json.Marshal(newBeverage)
	req := httptest.NewRequest(http.MethodPost, "/beverage", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос и проверяем результат
	resp, _ := app.Test(req)

	// Проверяем успешный статус
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
	mockStorage.AssertExpectations(t)
}

func TestCreateBeverage_Failure(t *testing.T) {
	// Создаем мок для StorageInterface
	mockStorage := new(mocks.StorageInterface)
	handler := NewHandler(mockStorage)

	// Подготовка данных напитка
	newBeverage := storage.Beverage{
		Name:        "Cola",
		Type:        "Soft Drink",
		Price:       1.99,
		Description: "A refreshing drink",
	}

	// Настраиваем мок: ожидаем, что CreateBeverage вернет ошибку
	mockStorage.On("CreateBeverage", mock.Anything, newBeverage).Return(assert.AnError)

	// Создаем Fiber приложение
	app := fiber.New()

	// Регистрируем хендлер
	app.Post("/beverage", handler.CreateBeverage)

	// Подготавливаем HTTP-запрос
	reqBody, _ := json.Marshal(newBeverage)
	req := httptest.NewRequest(http.MethodPost, "/beverage", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос и проверяем результат
	resp, _ := app.Test(req)

	// Проверяем, что был возвращен код 500 (Internal Server Error)
	assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	mockStorage.AssertExpectations(t)
}

func TestGetBeverage_Succses(t *testing.T) {
	mockStorage := new(mocks.StorageInterface)
	handler := NewHandler(mockStorage)

	newBeverage := storage.Beverage{
		Name:        "Cola",
		Type:        "Soft Drink",
		Price:       1.99,
		Description: "A refreshing drink",
	}
	id := 1

	mockStorage.On("GetBeverage", mock.Anything, id).Return(&newBeverage, nil)

	app := fiber.New()

	app.Get("/beverage/:id", handler.GetBeverage)

	target := fmt.Sprintf("/beverage/%v", id)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	mockStorage.AssertExpectations(t)
}

func TestGetBeverage_Failure(t *testing.T) {
	mockStorage := new(mocks.StorageInterface)
	handler := NewHandler(mockStorage)

	id := 1

	mockStorage.On("GetBeverage", mock.Anything, id).Return(nil, nil)

	app := fiber.New()

	app.Get("/beverage/:id", handler.GetBeverage)

	target := fmt.Sprintf("/beverage/%v", id)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	mockStorage.AssertExpectations(t)
}

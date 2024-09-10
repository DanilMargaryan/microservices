package controllers

import (
	"context"
	"github.com/DanilMargaryan/microservices/internal/storage"
	"github.com/gofiber/fiber/v3"
)

type MainController struct {
	storage storage.Storage
}

func NewController(stg storage.Storage) *MainController {
	return &MainController{storage: stg}
}

func (c *MainController) CreateBeverage(ctx fiber.Ctx) error {
	var newBeverage storage.Beverage

	if err := ctx.Bind().Body(&newBeverage); err != nil {
		return ctx.Status(400).SendString("Ошибка при парсинге тела запроса")
	}

	if err := c.storage.CreateBeverage(context.Background(), newBeverage); err != nil {
		return ctx.Status(500).SendString("Ошибка при добавлении напитка в базу данных")
	}

	return ctx.Status(201).SendString("Напиток успешно добавлен!")
}

func (c *MainController) GetAllBeverages(ctx fiber.Ctx) error {
	beverages, err := c.storage.GetAllBeverages(context.Background())
	if err != nil {
		return ctx.Status(500).SendString("Ошибка при получении данных")
	}
	return ctx.JSON(beverages)
}

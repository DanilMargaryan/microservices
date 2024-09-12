package servise

import (
	"context"
	"encoding/json"
	"github.com/DanilMargaryan/microservices/internal/dto"
	"github.com/DanilMargaryan/microservices/internal/storage"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"strconv"
)

type MainController struct {
	storage storage.Storage
}

func NewController(stg storage.Storage) *MainController {
	return &MainController{storage: stg}
}

func (c *MainController) CreateBeverage(ctx fiber.Ctx) error {
	var newBeverage storage.Beverage

	if err := json.Unmarshal(ctx.Body(), &newBeverage); err != nil {
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}

	if err := c.storage.CreateBeverage(context.Background(), newBeverage); err != nil {
		log.Error(err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
		Data:   "The drink has been added successfully!",
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *MainController) GetAllBeverages(ctx fiber.Ctx) error {
	beverages, err := c.storage.GetAllBeverages(context.Background())
	if err != nil {
		log.Error(err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
		Data:   beverages,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *MainController) GetBeverage(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, _ := strconv.Atoi(idStr)
	beverage, err := c.storage.GetBeverage(context.Background(), id)
	if err != nil {
		log.Error(err)
		return dto.InternalServerError(ctx)
	}

	if beverage == nil {
		return dto.BadResponseError(ctx, dto.FieldIncorrect, "Drink not found")
	}

	response := dto.Response{
		Status: "success",
		Data:   beverage,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

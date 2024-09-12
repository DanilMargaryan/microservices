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

type MainHandlerInterface interface {
	CreateBeverage(ctx fiber.Ctx) error
	GetAllBeverages(ctx fiber.Ctx) error
	GetBeverage(ctx fiber.Ctx) error
}

type MainHandler struct {
	storage storage.StorageInterface
}

func NewHandler(stg storage.StorageInterface) *MainHandler {
	return &MainHandler{storage: stg}
}

func (c *MainHandler) CreateBeverage(ctx fiber.Ctx) error {
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

func (c *MainHandler) GetAllBeverages(ctx fiber.Ctx) error {
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

func (c *MainHandler) GetBeverage(ctx fiber.Ctx) error {
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

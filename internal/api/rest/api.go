package rest

import (
	"github.com/DanilMargaryan/microservices/internal/servise"
	"github.com/DanilMargaryan/microservices/internal/storage"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, stg storage.Storage) {
	beverageController := servise.NewController(stg)

	app.Get("/beverages", beverageController.GetAllBeverages)
	app.Get("/beverage/:id", beverageController.GetBeverage)

	app.Post("/beverage", beverageController.CreateBeverage)
}

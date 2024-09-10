package rest

import (
	"github.com/DanilMargaryan/microservices/internal/api/rest/controllers"
	"github.com/DanilMargaryan/microservices/internal/storage"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, stg storage.Storage) {
	beverageController := controllers.NewController(stg)

	app.Get("/beverages", beverageController.GetAllBeverages)
	app.Get("/beverage/:id", beverageController.GetBeverage)

	app.Post("/beverage", beverageController.CreateBeverage)
}

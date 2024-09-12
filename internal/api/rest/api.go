package rest

import (
	"github.com/DanilMargaryan/microservices/internal/servise"
	"github.com/gofiber/fiber/v3"
)

type Routers struct {
	MainHandler servise.MainHandlerInterface
}

func SetupRoutes(r *Routers) *fiber.App {
	app := fiber.New()

	app.Get("/beverages", r.MainHandler.GetAllBeverages)
	app.Get("/beverage/:id", r.MainHandler.GetBeverage)

	app.Post("/beverage", r.MainHandler.CreateBeverage)

	return app
}

package main

import (
	_ "github.com/DanilMargaryan/microservices/docs"
	"github.com/DanilMargaryan/microservices/internal/api/rest"
	"github.com/DanilMargaryan/microservices/internal/config"
	"github.com/DanilMargaryan/microservices/internal/storage"
	"github.com/gofiber/fiber/v3"
	"log"
)

// @title Beverage API
// @version 1.0
// @description API для управления напитками.
// @host localhost:3000
// @BasePath /v2

// @contact.name API Support
// @contact.email support@swagger.io
func main() {
	cfg := config.MustLoad()

	stg, err := storage.New(cfg)

	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer stg.Close()

	app := fiber.New()

	rest.SetupRoutes(app, *stg)

	log.Fatal(app.Listen(":3000"))
}

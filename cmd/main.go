package main

import (
	"github.com/DanilMargaryan/microservices/internal/api/rest"
	"github.com/DanilMargaryan/microservices/internal/config"
	"github.com/DanilMargaryan/microservices/internal/storage"
	"github.com/gofiber/fiber/v3"
	"log"
)

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

package main

import (
	"context"
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

	app.Get("/", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, World 👋!")
	})

	app.Get("/beverages", func(ctx fiber.Ctx) error {
		beverages, err := stg.GetAllBeverages(context.Background())
		if err != nil {
			return ctx.Status(500).SendString("Ошибка при получении данных")
		}
		return ctx.JSON(beverages)
	})

	app.Post("/newbeverage", func(ctx fiber.Ctx) error {
		var newBeverage storage.Beverage

		if err := ctx.Bind().Body(&newBeverage); err != nil {
			return ctx.Status(400).SendString("Ошибка при парсинге тела запроса")
		}

		if err := stg.CreateBeverage(context.Background(), newBeverage); err != nil {
			return ctx.Status(500).SendString("Ошибка при добавлении напитка в базу данных")
		}

		return ctx.Status(201).SendString("Напиток успешно добавлен!")
	})

	log.Fatal(app.Listen(":3000"))
}

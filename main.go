package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"ip-tracker/database"
	"ip-tracker/handler"
	"ip-tracker/model/web"
	"ip-tracker/service"
)

func main() {
	db := database.NewDatabase()
	app := fiber.New(fiber.Config{
		AppName:      "Ip Tracker",
		ErrorHandler: handler.ErrorHandler,
	})

	// middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// TODO: add dependency injection

	// service
	locationService := service.NewLocationService(db)

	// handler
	handler.NewLocationHandler(locationService).Bind(app)

	// health check
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(web.CommonResponse[string]{
			Status: "Success",
			Data:   "Hello, World!",
		})
	})

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}

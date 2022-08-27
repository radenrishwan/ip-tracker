package handler

import (
	"github.com/gofiber/fiber/v2"
	"ip-tracker/exception"
	"ip-tracker/service"
)

type LocationHandler interface {
	Find(ctx *fiber.Ctx) error
	Bind(router *fiber.App)
}

type locationHandler struct {
	service.LocationService
}

func NewLocationHandler(locationService service.LocationService) LocationHandler {
	return &locationHandler{LocationService: locationService}
}

func (handler *locationHandler) Find(ctx *fiber.Ctx) error {
	ip := ctx.Query("ip", "")
	if ip == "" {
		panic(exception.NewValidationException("ip cannot be null"))
	}

	result := handler.LocationService.Find(ip)

	return ctx.JSON(result)
}

func (handler *locationHandler) Bind(router *fiber.App) {
	router.Get("/api/location", handler.Find)
}

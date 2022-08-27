package handler

import (
	"github.com/gofiber/fiber/v2"
	"ip-tracker/exception"
	"ip-tracker/model/web"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(exception.ValidationException)
	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(web.CommonResponse[string]{
			Status: "Validation error",
			Data:   err.Error(),
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(web.CommonResponse[string]{
		Status: "Internal Server Error",
		Data:   err.Error(),
	})
}

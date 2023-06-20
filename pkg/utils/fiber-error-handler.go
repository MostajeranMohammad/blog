package utils

import (
	"errors"

	"github.com/MostajeranMohammad/blog/internal/entity"
	"github.com/MostajeranMohammad/blog/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FiberErrorHandler(logger logger.Interface) func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(404).JSON(entity.ResponseModel{Successful: false, Message: "record not found"})
		}

		if code == fiber.StatusInternalServerError {
			logger.Error(err.Error())
			return ctx.Status(code).JSON(entity.ResponseModel{Successful: false, Message: "Internal Server Error"})
		}
		return ctx.Status(code).JSON(entity.ResponseModel{Successful: false, Message: err.Error()})
	}
}

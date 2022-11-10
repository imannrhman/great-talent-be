package exception

import (
	"github.com/gofiber/fiber/v2"
	"great-talent-be/model"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.Status(400).JSON(model.WebResponse{
			Code:     400,
			Status:   "BAD_REQUEST",
			Messages: err.Error(),
		})
	}
	return ctx.Status(500).JSON(model.WebResponse{
		Code:     500,
		Status:   "INTERNAL_SERVER_ERROR",
		Messages: err.Error(),
	})
}

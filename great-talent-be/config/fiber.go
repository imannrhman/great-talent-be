package config

import (
	"github.com/gofiber/fiber/v2"
	"great-talent-be/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}

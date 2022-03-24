package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"great-talent-be/config"
	"great-talent-be/controller"
	"great-talent-be/exception"
	"great-talent-be/repository"
	"great-talent-be/service"
)

func main() {
	configuration := config.New()
	fmt.Print(configuration.Get("MONGO_URI"))
	database := config.NewMongoDatabase(configuration)

	employeeRepository := repository.NewEmployeeRepository(database)

	employeeService := service.NewEmployeeService(&employeeRepository)

	employeeController := controller.NewEmployeeController(&employeeService)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	employeeController.Route(app)

	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}

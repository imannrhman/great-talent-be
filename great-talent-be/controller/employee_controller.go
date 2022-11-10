package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"great-talent-be/exception"
	"great-talent-be/model"
	"great-talent-be/service"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService *service.EmployeeService) EmployeeController {
	return EmployeeController{
		EmployeeService: *employeeService,
	}
}

func (controller *EmployeeController) Route(app *fiber.App) {
	app.Post("/api/employee", controller.Create)
	app.Get("/api/employees", controller.List)
	app.Get("/api/total-salary/:id", controller.TotalSalary)
	app.Put("/api/employee/:id", controller.Create)
	app.Patch("/api/employee/:id", controller.Create)
	app.Delete("/api/employee/:id", controller.Create)
}

func (controller *EmployeeController) Create(ctx *fiber.Ctx) error {
	var request model.Employee
	err := ctx.BodyParser(&request)
	request.ID = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.EmployeeService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *EmployeeController) List(ctx *fiber.Ctx) error {
	responses := controller.EmployeeService.List()
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *EmployeeController) TotalSalary(ctx *fiber.Ctx) error {
	response := controller.EmployeeService.TotalSalary(ctx.Params("id"))
	return ctx.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

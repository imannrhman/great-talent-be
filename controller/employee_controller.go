package controller

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"great-talent-be/exception"
	"great-talent-be/model"
	"great-talent-be/service"
	"html/template"
)

var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService *service.EmployeeService) EmployeeController {
	return EmployeeController{
		EmployeeService: *employeeService,
	}
}

func (controller *EmployeeController) Route(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON("Welcome")
	})
	app.Post("/api/employee", controller.Create)
	app.Patch("/api/employee/:id", controller.Update)
	app.Delete("/api/employee/:id", controller.Delete)
	app.Get("/api/employees", controller.List)
	app.Get("/api/total-salary/:id", controller.TotalSalary)
}

func (controller *EmployeeController) Create(ctx *fiber.Ctx) error {
	var request model.Employee
	err := ctx.BodyParser(&request)
	request.ID = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.EmployeeService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:     201,
		Status:   "OK",
		Messages: "Berhasil mendaftarkan karyawan !",
		Data:     response,
	})
}

func (controller *EmployeeController) Update(ctx *fiber.Ctx) error {
	var request model.Employee
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)
	response := controller.EmployeeService.Update(ctx.Params("id"), request)
	return ctx.JSON(model.WebResponse{
		Code:     200,
		Status:   "OK",
		Messages: "Berhasil mengupdate karyawan !",
		Data:     response,
	})
}

func (controller *EmployeeController) Delete(ctx *fiber.Ctx) error {
	response := controller.EmployeeService.Delete(ctx.Params("id"))
	return ctx.JSON(model.WebResponse{
		Code:     200,
		Messages: "Berhasil menghapus karyawan !",
		Data:     response,
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

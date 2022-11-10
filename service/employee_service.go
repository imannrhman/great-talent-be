package service

import "great-talent-be/model"

type EmployeeService interface {
	Create(request model.Employee) (response model.Employee)
	Update(id string, request model.Employee) (response model.Employee)
	Delete(id string) (response string)
	List() (responses []model.Employee)
	TotalSalary(id string) (response model.TotalSalary)
}

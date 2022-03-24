package service

import "great-talent-be/model"

type EmployeeService interface {
	Create(request model.Employee) (response model.Employee)
	List() (responses []model.Employee)
	TotalSalary(id string) (response model.TotalSalary)
}

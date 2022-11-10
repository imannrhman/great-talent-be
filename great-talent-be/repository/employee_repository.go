package repository

import (
	"great-talent-be/entity"
)

type EmployeeRepository interface {
	Insert(employee entity.Employee)
	Update(employeeID string)
	Delete(employeeID string)
	FetchAll() (employees []entity.Employee)
	FindOne(employeeID string) entity.Employee
}

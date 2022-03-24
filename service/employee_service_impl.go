package service

import (
	"great-talent-be/entity"
	"great-talent-be/model"
	"great-talent-be/repository"
	"great-talent-be/validation"
)

func NewEmployeeService(employeeRepository *repository.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{
		*employeeRepository,
	}
}

type employeeServiceImpl struct {
	repository.EmployeeRepository
}

func (service employeeServiceImpl) Create(request model.Employee) (response model.Employee) {
	validation.Validate(request)

	employee := entity.Employee{
		ID:         request.ID,
		NIK:        request.NIK,
		Name:       request.Name,
		Class:      request.Class,
		Salary:     salary(request.Class),
		Gender:     request.Gender,
		Allowance:  request.Allowance,
		SalaryCuts: request.SalaryCuts,
	}

	service.EmployeeRepository.Insert(employee)

	response = model.Employee{
		ID:         request.ID,
		NIK:        request.NIK,
		Name:       request.Name,
		Class:      request.Class,
		Gender:     request.Gender,
		Salary:     salary(request.Class),
		Allowance:  request.Allowance,
		SalaryCuts: request.SalaryCuts,
	}
	return response
}

func (service employeeServiceImpl) List() (responses []model.Employee) {
	employees := service.EmployeeRepository.FetchAll()

	for _, employee := range employees {
		responses = append(responses, model.Employee{
			ID:         employee.ID,
			NIK:        employee.NIK,
			Name:       employee.Name,
			Class:      employee.Class,
			Gender:     employee.Gender,
			Salary:     employee.Salary,
			Allowance:  employee.Allowance,
			SalaryCuts: employee.SalaryCuts,
		})
	}

	return responses
}

func (service employeeServiceImpl) TotalSalary(id string) (response model.TotalSalary) {
	employee := service.EmployeeRepository.FindOne(id)

	totalSalary := (employee.Salary + employee.Allowance) - employee.SalaryCuts

	response = model.TotalSalary{
		Employee: model.Employee{
			ID:         employee.ID,
			NIK:        employee.NIK,
			Name:       employee.Name,
			Class:      employee.Class,
			Gender:     employee.Gender,
			Salary:     employee.Salary,
			Allowance:  employee.Allowance,
			SalaryCuts: employee.SalaryCuts,
		},
		TotalSalary: totalSalary,
	}

	return response
}

func salary(class int32) float64 {
	switch class {
	case 1:
		return 2686500 //Salary Class I
	case 2:
		return 3820000 //Salary Class II
	case 3:
		return 4797000 //Salary Class III
	case 4:
		return 5901200 //Salary Class VI
	default:
		return 0 //Undetected
	}
}

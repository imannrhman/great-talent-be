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

func (service employeeServiceImpl) Update(id string, request model.Employee) (response model.Employee) {
	employee := service.EmployeeRepository.FindOne(id)
	if request.NIK != "" {
		employee.NIK = request.NIK
	}

	if request.Name != "" {
		employee.Name = request.Name
	}

	if request.Gender != "" {
		employee.Gender = request.Gender
	}

	if request.Class != 0 {
		employee.Class = request.Class
		employee.Salary = salary(request.Class)
	}

	if request.Allowance != 0.0 {
		employee.Allowance = request.Allowance
	}

	if request.SalaryCuts != 0.0 {
		employee.SalaryCuts = request.SalaryCuts
	}

	validation.UpdateValidate(request)

	employee = service.EmployeeRepository.Update(id, employee)
	response = model.Employee{
		ID:         employee.ID,
		NIK:        employee.NIK,
		Name:       employee.Name,
		Class:      employee.Class,
		Gender:     employee.Gender,
		Salary:     employee.Salary,
		Allowance:  employee.Allowance,
		SalaryCuts: employee.SalaryCuts,
	}

	return response
}

func (service employeeServiceImpl) Delete(id string) (response string) {
	service.EmployeeRepository.Delete(id)
	return id
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

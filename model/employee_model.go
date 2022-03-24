package model

type Employee struct {
	ID         string  `json:"id"`
	NIK        string  `json:"nik"`
	Name       string  `json:"name"`
	Gender     string  `json:"gender"`
	Class      int32   `json:"class"`
	Salary     float64 `json:"salary"`
	Allowance  float64 `json:"allowance"`
	SalaryCuts float64 `json:"salary_cuts"`
}

type TotalSalary struct {
	Employee
	TotalSalary float64 `json:"total_salary"`
}

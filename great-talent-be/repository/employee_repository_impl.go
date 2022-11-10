package repository

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"great-talent-be/config"
	"great-talent-be/entity"
	"great-talent-be/exception"
	"math"
)

func NewEmployeeRepository(database *mongo.Database) EmployeeRepository {
	return &employeeRepositoryImpl{
		database.Collection("employees"),
	}
}

type employeeRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository employeeRepositoryImpl) Insert(employee entity.Employee) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":        employee.ID,
		"nik":        employee.NIK,
		"name":       employee.Name,
		"gender":     employee.Gender,
		"class":      employee.Class,
		"salary":     employee.Salary,
		"allowance":  employee.Allowance,
		"salary_cut": employee.SalaryCuts,
	})
	exception.PanicIfNeeded(err)
}

func (repository employeeRepositoryImpl) FetchAll() (employees []entity.Employee) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)

	exception.PanicIfNeeded(err)

	for _, document := range documents {
		employee := entity.Employee{}
		employee.ID = document["_id"].(string)
		employee.NIK = document["nik"].(string)
		employee.Name = document["name"].(string)
		employee.Gender = document["gender"].(string)

		employee.Class = document["class"].(int32)
		employee.Salary = math.Round(document["salary"].(float64))
		employee.Allowance = math.Round(document["allowance"].(float64))
		employee.SalaryCuts = math.Round(document["salary_cut"].(float64))

		employees = append(employees, employee)
	}
	return employees
}

func (repository employeeRepositoryImpl) FindOne(employeeID string) entity.Employee {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var document bson.M
	err := repository.Collection.FindOne(ctx, bson.M{"_id": employeeID}).Decode(&document)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	employee := entity.Employee{}
	employee.ID = document["_id"].(string)
	employee.NIK = document["nik"].(string)
	employee.Name = document["name"].(string)
	employee.Gender = document["gender"].(string)

	employee.Class = document["class"].(int32)
	employee.Salary = math.Round(document["salary"].(float64))
	employee.Allowance = math.Round(document["allowance"].(float64))
	employee.SalaryCuts = math.Round(document["salary_cut"].(float64))

	return employee
}

func (repository employeeRepositoryImpl) Update(employeeID string) {

}

func (repository employeeRepositoryImpl) Delete(employeeID string) {

}

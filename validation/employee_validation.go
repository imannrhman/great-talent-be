package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"great-talent-be/exception"
	"great-talent-be/model"
)

func Validate(request model.Employee) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.ID, validation.Required),
		validation.Field(&request.NIK, validation.Required, validation.RuneLength(16, 16)),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Gender, validation.Required, validation.In("Male", "Female")),
		validation.Field(&request.Class, validation.Required, validation.Max(4)),
		validation.Field(&request.Allowance),
		validation.Field(&request.SalaryCuts),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

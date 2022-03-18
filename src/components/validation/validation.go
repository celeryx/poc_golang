package validation

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(model interface{}) []*ErrorResponse {

	var errors []*ErrorResponse
	err := validate.Struct(model)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			errNamespace := strings.Split(err.StructNamespace(), ".")[1:]
			element.FailedField = strings.Join(errNamespace, ".")
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

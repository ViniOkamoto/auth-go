package validator

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type ValidateError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateStruct(s interface{}) json.RawMessage {
	if validate == nil {
		validate = validator.New()
	}

	err := validate.Struct(s)

	if err == nil {
		return nil
	}

	var errors []ValidateError

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, ValidateError{
			Field:   err.Field(),
			Message: err.Tag(),
		})
	}

	result, _ := json.Marshal(errors)

	return result
}

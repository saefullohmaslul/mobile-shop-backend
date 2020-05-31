package validations

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/saefullohmaslul/mobile-shop-backend/src/apps/libraries/response"
)

// Validate is method to validate error schema
func Validate(schema interface{}, errors []response.Error) {
	validate := validator.New()

	if err := validate.Struct(schema); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errors = append(errors, response.Error{
				Message: err.Error(),
				Flag:    "INVALID_VALIDATION_SCHEMA",
			})
		}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, response.Error{
				Message: fmt.Sprint(err),
				Flag:    fmt.Sprintf("INVALID_VALIDATION_%s", strings.ToUpper(err.Field())),
			})
		}

		response.BadRequest("Validation error", errors)
	}

	if errors != nil {
		response.BadRequest("Validation error", errors)
	}
}

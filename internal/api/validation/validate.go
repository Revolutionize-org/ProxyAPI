package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string
	Tag   string
}

func (e *ValidationError) Format() map[string]string {
	return map[string]string{
		e.Field: fmt.Sprintf("A %s is %s", e.Field, e.Tag),
	}
}

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func Validate(data interface{}) []ValidationError {
	var errors []ValidationError
	err := validate.Struct(data)

	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			element := ValidationError{
				Field: strings.ToLower(vErr.Field()),
				Tag:   strings.ToLower(vErr.Tag()),
			}
			errors = append(errors, element)
		}
	}
	return errors
}

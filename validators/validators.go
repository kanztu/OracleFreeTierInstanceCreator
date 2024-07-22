package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Fields         []string
	ValidationTags []string
}

type StandaloneValidationError struct {
	Field         string
	ValidationTag string
}

// Error returns the first validation error out of all errors
func (v ValidationError) Error() string {
	errStr := "Failed validation on fields: \n"
	for i, field := range v.Fields {
		errStr += fmt.Sprintf("Field %v failed on condition %v\n", field, v.ValidationTags[i])
	}
	return errStr
}

func Validate(data interface{}) error {
	v := validator.New()

	err := v.Struct(data)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	var accumulatedErrors = ValidationError{}
	for _, validationErr := range validationErrors {
		accumulatedErrors.Fields = append(accumulatedErrors.Fields, validationErr.Field())
		accumulatedErrors.ValidationTags = append(accumulatedErrors.ValidationTags, validationErr.Tag())
	}
	return accumulatedErrors
}

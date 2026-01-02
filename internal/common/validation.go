package common

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationErrors converts validator errors into a user-friendly map
func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrs {
			fieldName := toSnakeCase(fieldError.Field())
			errors[fieldName] = getErrorMessage(fieldError)
		}
	}

	return errors
}

// getErrorMessage returns a user-friendly error message for a validation error
func getErrorMessage(fieldError validator.FieldError) string {
	field := fieldError.Field()
	tag := fieldError.Tag()
	param := fieldError.Param()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		if fieldError.Type().String() == "string" {
			return fmt.Sprintf("%s must be at least %s characters long", field, param)
		}
		return fmt.Sprintf("%s must be at least %s", field, param)
	case "max":
		if fieldError.Type().String() == "string" {
			return fmt.Sprintf("%s must not exceed %s characters", field, param)
		}
		return fmt.Sprintf("%s must not exceed %s", field, param)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", field, param)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, param)
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, param)
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, param)
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, param)
	case "alpha":
		return fmt.Sprintf("%s must contain only alphabetic characters", field)
	case "alphanum":
		return fmt.Sprintf("%s must contain only alphanumeric characters", field)
	case "numeric":
		return fmt.Sprintf("%s must be a valid number", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "uri":
		return fmt.Sprintf("%s must be a valid URI", field)
	case "oneof":
		return fmt.Sprintf("%s must be one of [%s]", field, param)
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", field)
	case "eqfield":
		return fmt.Sprintf("%s must equal %s", field, param)
	case "nefield":
		return fmt.Sprintf("%s must not equal %s", field, param)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// toSnakeCase converts a string from camelCase or PascalCase to snake_case
func toSnakeCase(str string) string {
	var result strings.Builder
	for i, char := range str {
		if i > 0 && char >= 'A' && char <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(char)
	}
	return strings.ToLower(result.String())
}

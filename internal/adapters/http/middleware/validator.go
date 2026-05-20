package middleware

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func TranslateValidationErrors(err error) string {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				return fmt.Sprintf("el campo '%s' es obligatorio", e.Field())
			case "min":
				return fmt.Sprintf("el campo '%s' debe tener mínimo %s caracteres", e.Field(), e.Param())
			case "max":
				return fmt.Sprintf("el campo '%s' debe tener máximo %s caracteres", e.Field(), e.Param())
			case "gt":
				return fmt.Sprintf("el campo '%s' debe ser mayor a %s", e.Field(), e.Param())
			}
		}
	}
	return "datos inválidos"
}

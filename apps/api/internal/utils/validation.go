package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, strings.ToLower(e.Field())+" "+e.Tag())
	}

	return  errors
}

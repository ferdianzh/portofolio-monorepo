package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()

	Validate.RegisterValidation(
		"strong_password",
		PasswordValidation,
	)
}

func ValidateStruct(data interface{}) []string {
	if err := Validate.Struct(data); err != nil {
		return FormatValidationErrors(err)
	}

	return nil
}

func FormatValidationErrors(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
			case "strong_password":
				errors = append(errors, "Password must contain letters, numbers, and symbols")
			default:
				errors = append(errors, e.Field()+" "+e.Tag())
		}
	}

	return errors
}

var (
	hasLetter = regexp.MustCompile(`[A-Za-z]`)
	hasNumber = regexp.MustCompile(`[0-9]`)
	hasSymbol = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?]`)
)

func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	return hasLetter.MatchString(password) &&
		hasNumber.MatchString(password) &&
		hasSymbol.MatchString(password)
}

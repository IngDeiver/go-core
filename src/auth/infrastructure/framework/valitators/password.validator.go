package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
    validate = validator.New()
    validate.RegisterValidation("password", ValidatePassword)
}

func ValidatePassword(fl validator.FieldLevel) bool {
    password := fl.Field().String()
    hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
    hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
    //hasSpecial := regexp.MustCompile(`[!@#~$%^&*()_+\-|{}[:;],.?<>]`).MatchString(password)
    return hasLetter  && hasNumber
}

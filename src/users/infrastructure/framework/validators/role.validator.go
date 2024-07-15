package validations

import (
	"github.com/go-playground/validator/v10"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("role", ValidateRole)
}

// validate valid values to use in role property
func ValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	if role == "" {
		return true // Allow empty role
	}

	switch userDomain.Role(role) {
	case userDomain.AdminRole, userDomain.UserRole, userDomain.ManagerRole:
		return true
	}
	return false
}

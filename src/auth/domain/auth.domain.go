package authDomain

import (
	"github.com/golang-jwt/jwt/v5"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
)

type Auth struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func New(user userDomain.User, token string) Auth {
	return Auth{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}
}

type AuthClaims struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
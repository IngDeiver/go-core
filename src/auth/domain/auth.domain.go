package authDomain

import (
	"github.com/golang-jwt/jwt/v5"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
)

type AuthWithToken struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type Auth struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func NewAuthTokenResponse(user userDomain.User, token string) AuthWithToken {
	return AuthWithToken{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}
}

type AuthWithClaims struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
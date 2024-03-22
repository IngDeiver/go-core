package authService

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authDomain "github.com/ingdeiver/go-core/src/auth/domain"
	authDto "github.com/ingdeiver/go-core/src/auth/domain/dto"
	errDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userRepo "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
)

type AuthService struct {
	userRepository  *userRepo.UserRepository
}

func (service *AuthService) Login(login authDto.LoginDto) (authDomain.Auth, error) {
	
	var user *userDomain.User
	var response authDomain.Auth

	//validate if exist by email

	if user == nil {
		return response, errDomain.ErrUnauthorizedError
	}

	//validate password
	if user.Password != login.Password {
		return response, errDomain.ErrUnauthorizedError
	}

	token, err := createToken(*user)
	if err != nil {
		return response, errDomain.ErrInternalServerError
	}

	response = authDomain.New(*user, token)
	return response, nil
}

func createToken (user userDomain.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := authDomain.AuthClaims{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return token, err
}

/*func validateToken(tokenString string) (bool, error){
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	
	switch {
	case token.Valid:
		return true, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		err := errors.New("That's not even a token")
		log.Println(err)
		return false, err
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		err := errors.New("Invalid signature")
		log.Println(err)
		return false, err
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		// Invalid signature
		err := errors.New("Timing is everything")
		log.Println(err)
		return false, err
	default:
		log.Println("Couldn't handle this token:", err)
		return false, err
	}
}*/

func New(repo  *userRepo.UserRepository) AuthService{
return AuthService{repo}
}
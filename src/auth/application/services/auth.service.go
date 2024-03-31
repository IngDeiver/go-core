package authService

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authDomain "github.com/ingdeiver/go-core/src/auth/domain"
	authDto "github.com/ingdeiver/go-core/src/auth/domain/dto"
	errDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
	email "github.com/ingdeiver/go-core/src/emails/application/services"

	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userRepo "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
)

var l = logger.Get()

type AuthService struct {
	userRepository  *userRepo.UserRepository
	emailService *email.EmailService
}

func New(repo  *userRepo.UserRepository, emailService *email.EmailService) *AuthService{
	return &AuthService{repo,emailService}
}

func (service *AuthService) Login(login authDto.LoginDto) (authDomain.AuthWithToken, error) {
	
	var user *userDomain.User
	var response authDomain.AuthWithToken

	//validate if exist by email

	user = &userDomain.User{"1","Deiver","Email", "PWD"}
	if user == nil {
		return response, errDomain.ErrUnauthorizedError
	}

	//validate password
	if user.Password != login.Password {
		return response, errDomain.ErrUnauthorizedError
	}

	token, err := createUserToken(*user)
	if err != nil {
		return response, err
	}

	response = authDomain.NewAuthTokenResponse(*user, token)
	return response, nil
}


func createUserToken (user userDomain.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := authDomain.AuthWithClaims{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func CreateGenericToken (body  map[string]interface{}, exp time.Duration) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := authDomain.GenericJWTClaims{
		Body: body,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func ValidateAuthToken(tokenString string) (*authDomain.Auth, error){
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &authDomain.AuthWithClaims{}, func (token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	
	switch {
	case token.Valid:
		if claims, ok := token.Claims.(*authDomain.AuthWithClaims); ok {
			return &authDomain.Auth{ID: claims.ID, Name: claims.Name, Email: claims.Email}, nil
		}
		return nil, errDomain.ErrUnauthorizedError
	case errors.Is(err, jwt.ErrTokenMalformed):
		l.Error().Msg("That's not even a token")
		return nil, err
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		l.Error().Msg("Invalid signature")
		return nil, err
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		// Invalid signature
		l.Error().Msg("Timing is everything")
		return nil, err
	default:
		l.Info().Msgf("Couldn't handle this token: %v", err.Error())
		return nil, err
	}
}

func ValidateGenericToken(tokenString string) (*authDomain.GenericJWTClaims, error){
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &authDomain.GenericJWTClaims{}, func (token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	
	switch {
	case token.Valid:
		if claims, ok := token.Claims.(*authDomain.GenericJWTClaims); ok {
			return claims, nil
		}
		return nil, errDomain.ErrUnauthorizedError
	case errors.Is(err, jwt.ErrTokenMalformed):
		l.Error().Msg("That's not even a token")
		return nil, err
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		l.Error().Msg("Invalid signature")
		return nil, err
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		// Invalid signature
		l.Error().Msg("Timing is everything")
		return nil, err
	default:
		l.Info().Msgf("Couldn't handle this token: %v", err.Error())
		return nil, err
	}
}


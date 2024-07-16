package authService

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authDomain "github.com/ingdeiver/go-core/src/auth/domain"
	authDto "github.com/ingdeiver/go-core/src/auth/domain/dto"
	errDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
	"github.com/ingdeiver/go-core/src/commons/infrastructure/helpers"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	emailInterfaces "github.com/ingdeiver/go-core/src/emails/domain"
	emailTypes "github.com/ingdeiver/go-core/src/emails/domain/constants"
	emailDomain "github.com/ingdeiver/go-core/src/emails/domain/interfaces"
	userServices "github.com/ingdeiver/go-core/src/users/application/services"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userDtos "github.com/ingdeiver/go-core/src/users/domain/dto"
	"go.mongodb.org/mongo-driver/bson"
)

var l = logger.Get()

type AuthService struct {
	userService *userServices.UserService
	emailService   emailDomain.EmailServiceDomain
}

func New(userService *userServices.UserService, emailService emailDomain.EmailServiceDomain) *AuthService {
	return &AuthService{userService, emailService}
}

func (service *AuthService) Login(login authDto.LoginDto) (authDomain.AuthWithToken, error) {
	var response authDomain.AuthWithToken
	//validate if exist by email
	userInfo := bson.M{"email": login.Email}
	user, err := service.userService.FindOne(userInfo)

	if user == nil || err != nil {
		return response, errDomain.ErrUnauthorizedError
	}

	//validate password
	if !helpers.CheckPasswordHash(login.Password, user.Password) {
		l.Warn().Msgf("%v no match password", user.FirstName)
		return response, errDomain.ErrUnauthorizedError
	}

	token, err := createUserToken(*user)
	if err != nil {
		return response, err
	}

	response = authDomain.NewAuthTokenResponse(*user, token)
	return response, nil
}

func (service *AuthService) SendWelcomeEmail(user userDomain.User) (bool, error) {
	body  := make(map[string]interface{})
	body["_id"] = user.ID
    body["action"] = "restore-password"

	token, err := CreateGenericToken(nil, 24 * time.Hour)

	if err != nil {
		return false, err
	}

	emailType := emailTypes.Password
	emailInfo := emailInterfaces.EmailMessageDomain{To: []string{user.Email}, Subject: "Welcome to Corsox - Set Your Password"}
	
	frontendURL := os.Getenv("FRONTEND_URL")
    if frontendURL == "" {
        return false, errors.New("FRONTEND_URL is not set")
    }

	template := emailInterfaces.EmailTemplateBodyDomain{
		Title: fmt.Sprintf("Dear %v, Welcome to %v!", user.FirstName, "[Your Application Name]"), 
		Message: "We are excited to have you join our community. To get started, please set your password by clicking the link below:",
		ButtomMessage: "Set password",
		ButtomURL: fmt.Sprintf("%v/%v/%v",frontendURL,"set-password",token),
	}

	return service.emailService.SendEmail(emailType, emailInfo, template)
}

func (service *AuthService) SendForgotEmail(user userDomain.User) (bool, error) {
	body  := make(map[string]interface{})
	body["_id"] = user.ID
    body["action"] = "restore-password"

	token, err := CreateGenericToken(nil, 24 * time.Hour)

	if err != nil {
		return false, err
	}

	emailType := emailTypes.Password
	emailInfo := emailInterfaces.EmailMessageDomain{To: []string{user.Email}, Subject: "Welcome to Corsox - Set Your Password"}
	
	frontendURL := os.Getenv("FRONTEND_URL")
    if frontendURL == "" {
        return false, errors.New("FRONTEND_URL is not set")
    }

	template := emailInterfaces.EmailTemplateBodyDomain{
		Title: fmt.Sprintf("Dear %v", user.FirstName), 
		Message: "We received a request to reset your password for your [Your Application Name] account. To reset your password, please click the link below:",
		ButtomMessage: "Restore password",
		ButtomURL: fmt.Sprintf("%v/%v/%v",frontendURL,"set-password",token),
	}

	return service.emailService.SendEmail(emailType, emailInfo, template)
}

func (service *AuthService) Register(register authDto.RegisterDto) (*authDomain.AuthWithToken, error) {
	var response authDomain.AuthWithToken
	user, err := service.userService.Create(userDtos.CreateUserDto{
		FirstName: register.FirstName,
		LastName: register.LastName,
		Email: register.Email,
		Password: register.Password,
		Role: register.Role,
	})
	if err != nil {
		return nil, err
	}
	token, err := createUserToken(user)
	if err != nil {
		return nil, err
	}

	response = authDomain.NewAuthTokenResponse(user, token)

	// send email
	_, err = service.SendWelcomeEmail(user)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func createUserToken(user userDomain.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	claims := authDomain.AuthWithClaims{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func CreateGenericToken(body map[string]interface{}, exp time.Duration) (string, error) {
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

func ValidateAuthToken(tokenString string) (*authDomain.Auth, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &authDomain.AuthWithClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	switch {
	case token.Valid:
		if claims, ok := token.Claims.(*authDomain.AuthWithClaims); ok {
			return &authDomain.Auth{ID: claims.ID, FirstName: claims.FirstName,
				LastName: claims.LastName, Email: claims.Email}, nil
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

func ValidateGenericToken(tokenString string) (*authDomain.GenericJWTClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(tokenString, &authDomain.GenericJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
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

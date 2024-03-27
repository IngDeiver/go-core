package authController

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	authService "github.com/ingdeiver/go-core/src/auth/application/services"
	authDto "github.com/ingdeiver/go-core/src/auth/domain/dto"
)


type AuthController struct {
	authService *authService.AuthService
}

func (s *AuthController) Login(c *gin.Context) {
	var loginDto authDto.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.Error(err)
		return
	}
	auth, err := s.authService.Login(loginDto)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK , auth)

}

func (s *AuthController) Some(c *gin.Context) {
	panic(errors.New("Testing recovery"))
	c.Status(http.StatusOK)
}

func New(s *authService.AuthService) AuthController{
	return AuthController{s}
}
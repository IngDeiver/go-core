package authController

import (
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

func (s *AuthController) Register(c *gin.Context) {
	var body authDto.RegisterDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}
	result, err := s.authService.Register(body)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated , result)

}

func New(s *authService.AuthService) *AuthController{
	return &AuthController{s}
}
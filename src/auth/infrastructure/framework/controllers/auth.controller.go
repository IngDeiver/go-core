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

func (s *AuthController) ForgotPassword(c *gin.Context) {
	var body authDto.ForgotPassword
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}
	err := s.authService.ForgotPassword(body)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK ,gin.H{"message": "Forgot email sent"})

}

func (s *AuthController) RestorePassword(c *gin.Context) {
	var body authDto.RestorePasswordDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}

	token := c.Param("token")
	err := s.authService.RestorePassword(token, body)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK ,gin.H{"message": "Password updated!"})

}

func New(s *authService.AuthService) *AuthController{
	return &AuthController{s}
}
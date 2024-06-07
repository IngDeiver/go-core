package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	userServices "github.com/ingdeiver/go-core/src/users/application/services"
)

type UserController struct {
	userService *userServices.UserService
}

func New (userService *userServices.UserService) *UserController{
	return &UserController{userService}
}

func (s *UserController) List(c *gin.Context) {
	response, err := s.userService.Base.List(nil)
	if err != nil {
		c.Error(err)
		return
	}
	fmt.Println(response)
	c.JSON(http.StatusOK , response)
}
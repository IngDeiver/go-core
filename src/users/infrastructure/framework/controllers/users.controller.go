package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	userServices "github.com/ingdeiver/go-core/src/users/application/services"
	userDtos "github.com/ingdeiver/go-core/src/users/domain/dto"
)

type UserController struct {
	userService *userServices.UserService
}

func New (userService *userServices.UserService) *UserController{
	return &UserController{userService}
}

func (s *UserController) List(c *gin.Context) {
	var filterDTO userDtos.UserFilterDto
	if err := c.ShouldBindQuery(&filterDTO); err != nil {
		c.Error(err)
		return
	}

	var paginationDTO dtos.PaginationParamsDto
	if err := c.ShouldBindQuery(&paginationDTO); err != nil {
		c.Error(err)
		return
	}

	
	var sortDTO dtos.SortParamsDto
	if err := c.ShouldBindQuery(&sortDTO); err != nil {
		c.Error(err)
		return
	}

	response, err := s.userService.List(filterDTO, &paginationDTO, &sortDTO)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK , response)
}
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

	response, err := s.userService.FindAll(filterDTO, &paginationDTO, &sortDTO)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK , response)
}

func (s *UserController) All(c *gin.Context) {
	var filterDTO userDtos.UserFilterDto
	if err := c.ShouldBindQuery(&filterDTO); err != nil {
		c.Error(err)
		return
	}

	response, err := s.userService.FindAllWithoutPagination(filterDTO)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK , response)
}

func (s *UserController) Create(c *gin.Context) {
	var body userDtos.CreateUserDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}
	result, err := s.userService.Create(body)

	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusCreated , result)

}

func (s *UserController) UpdateById(c *gin.Context) {
	id := c.Param("id")
	var body userDtos.UpdateUserDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.Error(err)
		return
	}
	result, err := s.userService.UpdateById(id, body)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated , result)

}

func (s *UserController) FindById(c *gin.Context){
	id := c.Param("id")
	result, err := s.userService.FindById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK , result)
}

func (s *UserController) RemoveById(c *gin.Context){
	id := c.Param("id")
	result, err := s.userService.RemoveById(id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK , result)
}

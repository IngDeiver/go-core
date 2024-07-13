package userService

import (
	baseService "github.com/ingdeiver/go-core/src/commons/application/services/base"
	"github.com/ingdeiver/go-core/src/commons/domain/dtos"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userRepo "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
)

// composition from base service domain adb implements BaseServiceDomain
type UserService struct {
	base *baseService.BaseService[userDomain.User]
	//add another compositions here
}

func New(repository *userRepo.UserRepository ) *UserService{
	BaseService := baseService.New[userDomain.User](repository)
	return &UserService{base: &BaseService }
}

func (s *UserService) List(filter interface{}, pagination *dtos.PaginationParamsDto, sort *dtos.SortParamsDto) (*dtos.PagedResponse[userDomain.User], error) {
    return s.base.List(filter, pagination, sort)
}



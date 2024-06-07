package userService

import (
	baseService "github.com/ingdeiver/go-core/src/commons/application/services/base"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userRepo "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
)

// composition from base service domain adb implements BaseServiceDomain
type UserService struct {
	Base *baseService.BaseService[userDomain.User]
	//add another compositions here
}

func New(repository *userRepo.UserRepository ) *UserService{
	BaseService := baseService.New[userDomain.User](repository)
	return &UserService{Base: &BaseService }
}




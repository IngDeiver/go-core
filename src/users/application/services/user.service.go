package userService

import (
	baseMongoService "github.com/ingdeiver/go-core/src/commons/application/services/mongo"
	userDomain "github.com/ingdeiver/go-core/src/users/domain"
	userRepo "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
)

// composition from base service domain
type UserService struct {
	Base *baseMongoService.BaseMongoService[userDomain.User]
	//add another compositions here
}

func New(repository userRepo.UserRepository ) UserService{
	BaseService := baseMongoService.New[userDomain.User]()
	return UserService{Base: &BaseService }
}




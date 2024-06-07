package users

import (
	middlewares "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/middlewares"
	"github.com/ingdeiver/go-core/src/config"
	userServices "github.com/ingdeiver/go-core/src/users/application/services"
	userControllers "github.com/ingdeiver/go-core/src/users/infrastructure/framework/controllers"
	userRepositories "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
)


var UserRepository *userRepositories.UserRepository
var UserService *userServices.UserService
var UserController *userControllers.UserController

// Instance Repositories, Services, Controllers and more
func InitUsersModule(){
	router := config.GetRouter()

	// ----- repositories -----
	UserRepository = userRepositories.New()


	// ----- services -----
	UserService = userServices.New(UserRepository)

	// ----- controllers -----
	UserController = userControllers.New(UserService)

	userRouter := router.Group("/users")
	{
		userRouter.Use(middlewares.AuthMiddleware())
		userRouter.GET("", UserController.List)
	}
}
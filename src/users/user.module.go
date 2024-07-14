package users

import (
	middlewares "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/middlewares"
	mongoMiddleware "github.com/ingdeiver/go-core/src/commons/infrastructure/middlewares/mongo"
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
		userRouter.POST("", UserController.Create)
		userRouter.GET("", UserController.List)
		userRouter.GET("all", UserController.All)
		userRouter.GET(":id", mongoMiddleware.ValidateObjectID(), UserController.FindById)
		userRouter.DELETE(":id", mongoMiddleware.ValidateObjectID(), UserController.RemoveById)
		userRouter.PUT(":id", mongoMiddleware.ValidateObjectID(), UserController.UpdateById)
	}
}
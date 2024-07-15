package users

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	middlewares "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/middlewares"
	mongoMiddleware "github.com/ingdeiver/go-core/src/commons/infrastructure/middlewares/mongo"
	"github.com/ingdeiver/go-core/src/config"
	userServices "github.com/ingdeiver/go-core/src/users/application/services"
	userControllers "github.com/ingdeiver/go-core/src/users/infrastructure/framework/controllers"
	validations "github.com/ingdeiver/go-core/src/users/infrastructure/framework/validators"
	userRepositories "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
)


var UserRepository *userRepositories.UserRepository
var UserService *userServices.UserService
var UserController *userControllers.UserController

// register custom validator here
func registerValidators(){
	// role validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role", validations.ValidateRole)
	}
}

// Instance Repositories, Services, Controllers and more
func InitUsersModule(){
	// set user validators
	registerValidators()

	// ----- repositories -----
	UserRepository = userRepositories.New()


	// ----- services -----
	UserService = userServices.New(UserRepository)

	// ----- controllers -----
	router := config.GetRouter()
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
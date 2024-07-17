package auth

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	auth "github.com/ingdeiver/go-core/src/auth/application/services"
	authController "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/controllers"
	validations "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/valitators"
	"github.com/ingdeiver/go-core/src/config"
	"github.com/ingdeiver/go-core/src/emails"
	"github.com/ingdeiver/go-core/src/users"
)

var AuthService *auth.AuthService
var AuthController *authController.AuthController

// register custom validators here
func registerValidators(){
	// password validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validations.ValidatePassword)
	}
}


// Instance Repositories, Services, Controllers and more
func InitAuthModule(){
	registerValidators()

    //services
    AuthService = auth.New(users.UserService,emails.EmailService) // UserService and EmailService already was initialize
	AuthController = authController.New(AuthService)
    
    // controllers
	router := config.GetRouter()
    authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", AuthController.Login)
		authRouter.POST("/register", AuthController.Register)
		authRouter.POST("/forgot-password", AuthController.ForgotPassword)
		authRouter.POST("/restore-password/:token", AuthController.RestorePassword)
	}
}
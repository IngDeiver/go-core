package auth

import (
	auth "github.com/ingdeiver/go-core/src/auth/application/services"
	authController "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/controllers"
	"github.com/ingdeiver/go-core/src/config"
	"github.com/ingdeiver/go-core/src/emails"
	"github.com/ingdeiver/go-core/src/users"
)

var AuthService *auth.AuthService
var AuthController *authController.AuthController

// Instance Repositories, Services, Controllers and more
func InitAuthModule(){
    router := config.GetRouter()

    //services
    AuthService = auth.New(users.UserRepository,emails.EmailService) // UserRepositoy and EmailService already was initialize
	AuthController = authController.New(AuthService)
    
    // controllers
    authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", AuthController.Login)
	}
}
package main

import (
	"net/http"
	"os"
	"time"

	auth "github.com/ingdeiver/go-core/src/auth/application/services"
	authControl "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/controllers"
	authMiddleware "github.com/ingdeiver/go-core/src/auth/infrastructure/framework/middlewares"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	email "github.com/ingdeiver/go-core/src/emails/application/services"
	smtp "github.com/ingdeiver/go-core/src/emails/infrastructure/gomail"
	httpServer "github.com/ingdeiver/go-core/src/http-server/infrastructure"
	userRepositories "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
	wsDomain "github.com/ingdeiver/go-core/src/ws/domain"
	wsHandlers "github.com/ingdeiver/go-core/src/ws/infrastructure/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var l = logger.Get()


func main() {
	loadEnv()
	start()
}

func loadEnv() {
	env := os.Getenv("APP_ENV")
    var err error

    switch env {
    case "production":
        err = godotenv.Load(".env.production")
    case "development":
        err = godotenv.Load(".env.development")
    default:
        err = godotenv.Load(".env.local")
    }

    if err != nil {
        l.Fatal().Msg("Error loading .env file")
    }

	if len(env) > 0 {
		l.Info().Msgf("Environment loaded: %s", env)
	}else {
		l.Info().Msg("Environment loaded: local")
	}
}


/*func stop(){
	// close db coneection
	// close emails chanel
}*/

func start(){
	// ------------ create router ------------
	router := gin.New()
	
	s := &http.Server{
		//Addr:           ":8080", // Optional, the value of the PORT variable will be used if it exists, otherwise it will be used ,the pure 80
		Handler:        router, // Optional, pass null if you do not use a handler
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
	server := httpServer.New(s, port)
	
	
	// ------------ ws config ------------
	webSocketDomain := wsDomain.New()
	webSocketManager := wsHandlers.New(webSocketDomain)
	server.SetWebSocketHandler(webSocketManager.Handler(), router)


	//------------ static files ------------
	server.ConfigureStaticFiles("public", router) 
	
	// ------------ set middlewares ------------
	server.ConfigGlobalMiddlewares(router)

	// ------------ routes ------------
	// ---- pulic routes
	userRepository := userRepositories.New()
	authService := auth.New(&userRepository)
	authController := authControl.New(&authService)
	authRouter :=router.Group("/auth")
	{
		authRouter.POST("/login", authController.Login)
	}

	// ---- protected routes
	userRouter :=router.Group("/users")
	{
		userRouter.Use(authMiddleware.AuthMiddleware())
		userRouter.POST("/",  authController.Some)
	}
	

	// ---------- email service ------------
	smtpService := smtp.New()
	emailService := email.New(&smtpService)
	emailService.CreateEmailsDeamon() // start emails chanel
	/*emailsChannel <- &emailDomain.EmailChanel{EmailType: 1, 
		Message: emailDomain.EmailMessageDomain{To: []string{"ingendeiver@gmail.com"}, 
		Cc: nil, Subject: "Pruebas xdd"}, 
		TemplateBody: emailDomain.EmailTemplateBodyDomain{
			Title: "Este es el titulo",
			Message: "Este es el mensaje", 
			ButtomMessage: "Boton!", ButtomURL: "https://outlook.office.com/mail/inbox/id/AAQkAGZkYzE1NjAxLWQ4OWYtNGZhYS1hODE4LTZjNzUyZjliYzhjZAAQAOYY%2FZ9qYZFMvM0PjzGoPp0%3D"} }
	*/
	// ------------ start server ------------
	l.Info().Msgf("Starting server on port: %v \n", port)
	server.StartServer()
}



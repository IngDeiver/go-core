package main

import (
	"log"
	"net/http"
	"os"
	"time"

	httpServer "github.com/ingdeiver/go-core/src/http-server/infrastructure"
	userServices "github.com/ingdeiver/go-core/src/users/application/services"
	userRepositories "github.com/ingdeiver/go-core/src/users/infrastructure/mongo/repositories"
	wsDomain "github.com/ingdeiver/go-core/src/ws/domain"
	wsHandlers "github.com/ingdeiver/go-core/src/ws/infrastructure/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	loadEnv()
	start()
}

func loadEnv() {
	env := os.Getenv("ENVIRONMENT")
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
        log.Fatal("Error loading .env file")
    }

	if len(env) > 0 {
		log.Printf("Environment loaded: %s", env)
	}else {
		log.Print("Environment loaded: local")
	}
}

func start(){
	router := gin.Default()

	s := &http.Server{
		//Addr:           ":8080", // Optional, the value of the PORT variable will be used if it exists, otherwise it will be used ,the pure 80
		Handler:        router, // Optional, pass null if you do not use a handler
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server := httpServer.New(s)
	

	userRepository := userRepositories.New()
	userService := userServices.New(userRepository)
	userService.Base.List()
	
	// ws config
	webSocketDomain := wsDomain.New()
	webSocketManager := wsHandlers.New(webSocketDomain)
	server.SetWebSocketHandler(webSocketManager.Handler(), router)


	//Use it if you need to server static files
	server.ConfigureStaticFiles("public", router) 
	
	
	// start server
	server.StartServer()
}



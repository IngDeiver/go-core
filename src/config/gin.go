package config

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	httpServer "github.com/ingdeiver/go-core/src/http/infrastructure"
)

// ------------ create router ------------
var router *gin.Engine
var server *httpServer.HttpServer

func CreateRouter(){
	router = gin.New()
}

func GetRouter() *gin.Engine {
	if router == nil {
		l.Fatal().Msg("No router found")
	}
	return router
}

func CreateServer(){
	if router == nil {
		l.Fatal().Msg("No router found")
	}

	var s = &http.Server{
		//Addr:           ":8080", // Optional, the value of the PORT variable will be used if it exists, otherwise it will be used ,the pure 80
		Handler:        router, // Optional, pass null if you do not use a handler
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server = httpServer.New(s, port)
}

func GetServer() *httpServer.HttpServer {
	if server == nil {
		l.Fatal().Msg("No server found")
	}
	return server
}


package httpServer

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	http *http.Server
}

func New(http *http.Server) *HttpServer {

	if PORT := os.Getenv("PORT"); len(PORT) > 0 {
		log.Println("Listening on port", PORT)
		http.Addr = fmt.Sprintf(":%v", PORT)
	}

	return &HttpServer{http}
}

func (server *HttpServer) ConfigureStaticFiles(path string, router *gin.Engine) {
	if router == nil {
		log.Fatal("Did not provide a router")
	}
	router.StaticFS(path, http.Dir(path))
	
	log.Printf("The %s directory was configured to serve static files.", path)
}

func (s *HttpServer) SetWebSocketHandler(handler func(*gin.Context), router *gin.Engine) {
	if router == nil {
		log.Fatal("Did not provide a router")
	}
	router.GET("/ws", handler)
	log.Println("Web socket atached with /ws prefix")
}

func (server *HttpServer) StartServer() {

	if server.http == nil {
		log.Fatal("Not exists server instance")
	}

	go func() {
		err := server.http.ListenAndServe()
		if err != nil {
			errorFormat := fmt.Errorf("start server error => %v", err)
			log.Fatal(errorFormat)
		}
	}()

	server.gracefulShutdown()

}

func (server *HttpServer) gracefulShutdown() {
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	wait := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	server.http.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

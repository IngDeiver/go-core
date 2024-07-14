package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"

	"github.com/gin-gonic/gin"
	errorMiddlware "github.com/ingdeiver/go-core/src/commons/infrastructure/middlewares/error"
	loggerMiddleware "github.com/ingdeiver/go-core/src/commons/infrastructure/middlewares/logger"
	recoveryMiddleware "github.com/ingdeiver/go-core/src/commons/infrastructure/middlewares/recovery"
)

var l = logger.Get()

type HttpServer struct {
	http *http.Server
}

func New(http *http.Server, port string) *HttpServer {
	http.Addr = fmt.Sprintf(":%v", port)
	return &HttpServer{http}
}

func (server *HttpServer) ConfigureStaticFiles(path string, router *gin.Engine) {
	if router == nil {
		l.Fatal().Msg("Did not provide a router")
		return
	}
	router.StaticFS(path, http.Dir(path))
	
	l.Info().Msgf("The %s directory was configured to serve static files.", path)
}

func (s *HttpServer) SetWebSocketHandler(handler func(*gin.Context), router *gin.Engine) {
	if router == nil {
		l.Fatal().Msg("Did not provide a router")
		return
	}
	router.GET("/ws", handler)
	l.Info().Msg("Web socket atached with /ws prefix")
}

func (server *HttpServer) StartServer() {
	
	if server.http == nil {
		l.Fatal().Msg("Not exists server instance")
	}

	go func() {
		err := server.http.ListenAndServe()
		if err != nil {
			l.Fatal().Msgf("start server error => %v", err)
		}
	}()
	l.Info().Msgf("Starting server on Addr %v \n", server.http.Addr)
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
	l.Info().Msg("shutting down server")
	os.Exit(0)
}


func (s *HttpServer) ConfigGlobalMiddlewares(router *gin.Engine) {
	router.Use(loggerMiddleware.LoggerMiddleware())
	router.Use(recoveryMiddleware.CustomRecoveryMiddleware())
	router.Use(errorMiddlware.ErrorHandlingMiddleware)

    corsOrigins := os.Getenv("CORS_ORIGIN")
    allowedOrigins := strings.Split(corsOrigins, ",")
    config := cors.Config{
        AllowOrigins: allowedOrigins,
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
    }
	l.Info().Msgf(`CORS origin allowed: %v`, allowedOrigins)
	router.Use(cors.New(config))
}
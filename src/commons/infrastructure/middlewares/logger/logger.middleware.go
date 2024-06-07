package loggerMiddlware

import (
	"time"

	"github.com/gin-gonic/gin"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
)

var l = logger.Get()

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()

        // before request
        path := c.Request.URL.Path
        method := c.Request.Method

        c.Next()

        // after request
        latency := time.Since(start)
        status := c.Writer.Status()
        userAgent := c.Request.Header.Get("User-Agent")

        user, existUser :=  c.Get("user")

        if !existUser {
            user = "guest"
        }

        if status >= 200 && status <= 399 {
            l.Info().
            Str("method", method).
            Str("path", path).
            Int("status", status).
            Str("user_agent", userAgent).
            Dur("latency", latency).
            Interface("user", user).
            Msg("request handling:")
        } else {
            l.Error().
            Str("method", method).
            Str("path", path).
            Int("status", status).
            Str("user_agent", userAgent).
            Dur("latency", latency).
            Interface("user", user).
            Msg("request handling:")
        }
        
    }
}

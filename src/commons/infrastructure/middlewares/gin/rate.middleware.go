package middlewares

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	rateLimitStr := os.Getenv("RATE_LIMIT")
    burstLimitStr := os.Getenv("BURST_LIMIT")

    rateLimit, err := strconv.ParseFloat(rateLimitStr, 64)
    if err != nil {
        l.Error().Msgf("Invalid RATE_LIMIT value: %s", rateLimitStr)
    }

    burstLimit, err := strconv.Atoi(burstLimitStr)
    if err != nil {
        l.Error().Msgf("Invalid BURST_LIMIT value: %s", burstLimitStr)
    }
    limiter := rate.NewLimiter(rate.Limit(rateLimit), burstLimit)

    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
            return
        }
        c.Next()
    }
}
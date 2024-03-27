package authMiddleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	authService "github.com/ingdeiver/go-core/src/auth/application/services"
	errorsDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, bearerSchema) {
			c.Error(errorsDomain.ErrUnauthorizedError)
			return
		}

		token := authHeader[len(bearerSchema):]

		auth, err := authService.ValidateToken(token)
		if auth == nil || err != nil {
			c.Error(errorsDomain.ErrUnauthorizedError)
			return
		}

		c.Set("user", auth)
		c.Next()
	}
}

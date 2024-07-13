package authMiddleware

import (
	"net/http"
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
			c.AbortWithStatusJSON(http.StatusUnauthorized,  gin.H{
				"message": errorsDomain.ErrUnauthorizedError.Error(),
			})
			return
		}

		token := authHeader[len(bearerSchema):]

		auth, err := authService.ValidateAuthToken(token)
		if auth == nil || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,  gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set("user", auth)
		c.Next()
	}
}

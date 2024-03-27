package recoveryMiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errorsDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
)

var l = logger.Get()


func CustomRecoveryMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                l.Error().Msgf("Recovery error: %v\n", err)

                // sent message to client
                c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                    "message": errorsDomain.ErrInternalServerError.Error(),
                })
            }
        }()
        c.Next()
    }
}

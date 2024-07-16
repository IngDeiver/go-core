package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	errorsDomain "github.com/ingdeiver/go-core/src/commons/domain/errors"
)

// Management each error pushed by c.Error()
func ErrorHandlingMiddleware(c *gin.Context) {
    c.Next()
    for _, ginErr := range c.Errors {
        err := ginErr.Err
        l.Error().Err(err).Send()

        if validationErrors, ok := err.(validator.ValidationErrors); ok {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
                "message": validationErrors.Error(),
            })
            return
        }
        switch err {
            case errorsDomain.ErrNotFoundError:
                c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
                    "message": err.Error(),
                })
            case errorsDomain.ErrUnauthorizedError:
                c.AbortWithStatusJSON(http.StatusUnauthorized,  gin.H{
                    "message": err.Error(),
                })
            case errorsDomain.ErrUserAlreadyExistsError:
                c.AbortWithStatusJSON(http.StatusBadRequest,  gin.H{
                    "message": err.Error(),
                })
            default :
                c.AbortWithStatusJSON(http.StatusInternalServerError,  gin.H{
                    "message": errorsDomain.ErrInternalServerError.Error(),
                })
        }
    }
}

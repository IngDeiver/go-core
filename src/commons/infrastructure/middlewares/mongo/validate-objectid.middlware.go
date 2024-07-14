package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Middleware to validate ObjectID
func ValidateObjectID() gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        if _, err := primitive.ObjectIDFromHex(id); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ObjectID"})
            c.Abort()
            return
        }
        c.Next()
    }
}
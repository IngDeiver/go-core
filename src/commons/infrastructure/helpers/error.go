package helpers

import "github.com/gin-gonic/gin"

/*
	Add error to c.Errors and then can capture using handle error middleware
*/
func AddException(c *gin.Context, err error) (*gin.Error){
	ginErr :=  c.Error(err)
	return ginErr
}
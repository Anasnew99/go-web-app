package utils

import "github.com/gin-gonic/gin"

func SendSuccessResponse(c *gin.Context, status int, data interface{}, messsage string) {
	c.JSON(status, gin.H{
		"message": messsage,
		"status":  true,
		"data":    data,
	})
}

func SendErrorResponse(c *gin.Context, status int, errorCode string, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"message":    message,
		"error_code": errorCode,
		"status":     false,
	})
}

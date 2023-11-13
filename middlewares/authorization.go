package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthAdminRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Accessing Protected Path")
		c.Next()
	}
}

package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.FullPath(), ctx.Request.Method)
		ctx.Next()
	}
}

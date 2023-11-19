package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("%v %v %v", ctx.Request.Method, ctx.Request.URL, ctx.Request.Proto)
		ctx.Next()
	}
}

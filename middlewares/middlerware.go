package middlewares

import "github.com/gin-gonic/gin"

func AddGeneralMiddlewares(r *gin.Engine) {
	r.Use(Logger())
	r.Use(gin.Recovery())
}

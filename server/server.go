package server

import (
	"anasnew99/server/chat_app/middlewares"
	"anasnew99/server/chat_app/routers"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	middlewares.AddGeneralMiddlewares(r)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	routers.AddRouters(r)
	return r
}

func StartServer() {
	r := GetRouter()

	r.Run() // listen and serve on

}

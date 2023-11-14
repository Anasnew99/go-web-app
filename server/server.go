package server

import (
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/middlewares"
	"anasnew99/server/chat_app/routers"
	"os"

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
	db.Connect(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_NAME"))
	defer db.Disconnect()
	r.Run() // listen and serve on

}

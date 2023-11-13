package server

import (
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/db"
	"anasnew99/server/chat_app/middlewares"
	"anasnew99/server/chat_app/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	r := GetRouter()
	// Connect to MongoDB
	db.Connect(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_NAME"))
	var adminUser = controllers.User{
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: os.Getenv("ADMIN_PASSWORD"),
	}
	controllers.DeleteUser(adminUser.Username)
	_, err := controllers.InsertUser(adminUser)
	if err != nil {
		log.Println("Admin user already exists")
	}
	defer db.Disconnect()
	r.Run() // listen and serve on

}

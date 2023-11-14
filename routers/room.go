package routers

import (
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/middlewares"
	"anasnew99/server/chat_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoomRouters(c *gin.RouterGroup) {
	c.POST("/", func(c *gin.Context) {
		var room models.AddRoomObject
		if err := c.ShouldBindJSON(&room); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		room.RoomOwner = controllers.Claims.GetUserFromRequest(c).Username
		if _, err := controllers.Room.AddRoom(room); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Room created successfully"})
	})

	c.POST("/:id/join", func(c *gin.Context) {
		var user = controllers.Claims.GetUserFromRequest(c)
		var data struct {
			Password string `json:"password"`
		}
		var roomID = c.Param("id")

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := controllers.Room.JoinRoom(user.Username, roomID, data.Password); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Joined room successfully"})

	})

	protectedRoomGroup := c.Group("/:id")
	protectedRoomGroup.Use(middlewares.AuthRoomRequest())
	protectedRoomGroup.GET("/", func(c *gin.Context) {
		var room = controllers.Claims.GetRoomFromRequest(c)
		c.JSON(http.StatusOK, gin.H{"room": room})
	})

	protectedRoomGroup.POST("/leave", func(c *gin.Context) {

	})

}

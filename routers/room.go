package routers

import (
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoomRouters(c *gin.RouterGroup) {
	c.POST("/", func(c *gin.Context) {
		var room controllers.AddRoomObject
		if err := c.ShouldBindJSON(&room); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		room.RoomOwner = controllers.GetUserFromRequest(c).Username
		if _, err := controllers.AddRoom(room); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Room created successfully"})
	})

	c.POST("/:id/join", func(c *gin.Context) {
		var user = controllers.GetUserFromRequest(c)
		var data struct {
			Password string `json:"password"`
		}
		var roomID = c.Param("id")

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := controllers.JoinRoom(user.Username, roomID, data.Password); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Joined room successfully"})

	})

	protectedRoomGroup := c.Group("/:id")
	protectedRoomGroup.Use(middlewares.AuthRoomRequest())
	protectedRoomGroup.GET("/", func(c *gin.Context) {
		var room = controllers.GetRoomFromRequest(c)
		c.JSON(http.StatusOK, gin.H{"room": room})
	})

	protectedRoomGroup.POST("/leave", func(c *gin.Context) {

	})

}

package routers

import (
	"anasnew99/server/chat_app/constants"
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/middlewares"
	"anasnew99/server/chat_app/models"
	"anasnew99/server/chat_app/utils"
	"anasnew99/server/chat_app/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRoomRouters(c *gin.RouterGroup) {
	c.POST("/", func(c *gin.Context) {
		var room models.AddRoomObject
		if err := c.ShouldBindJSON(&room); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, err.Error())
			return
		}
		room.RoomOwner = controllers.Claims.GetUserFromRequest(c).Username
		if _, err := controllers.Room.AddRoom(room); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, "Room already exists, please choose another id")
			return
		}

		utils.SendSuccessResponse(c, http.StatusCreated, nil, "Room created successfully")
	})

	c.POST("/:id/join", func(c *gin.Context) {
		var user = controllers.Claims.GetUserFromRequest(c)
		var data struct {
			Password string `json:"password"`
		}
		var roomID = c.Param("id")

		if err := c.ShouldBindJSON(&data); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, err.Error())
			return
		}
		if err := controllers.Room.JoinRoom(user.Username, roomID, data.Password); err != nil {
			utils.SendErrorResponse(c, http.StatusForbidden, constants.INVALID_DATA, "Wrong password")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, nil, "Joined room successfully")

	})

	protectedRoomGroup := c.Group("/:id")
	protectedRoomGroup.Use(middlewares.AuthRoomRequest())
	protectedRoomGroup.GET("/", func(c *gin.Context) {
		var room = controllers.Claims.GetRoomFromRequest(c)
		utils.SendSuccessResponse(c, http.StatusOK, room, "")
	})

	protectedRoomGroup.DELETE("/", func(c *gin.Context) {
		var room = controllers.Claims.GetRoomFromRequest(c)
		if !controllers.Room.IsUserOwnerOfTheRoom(controllers.Claims.GetUserFromRequest(c).Username, room.Id) {
			utils.SendErrorResponse(c, http.StatusUnauthorized, constants.UNAUTHORIZED, "You are not the owner of the room")
			return
		}
		if err := controllers.Room.DeleteRoom(room.Id); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, "Room does not exist")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, nil, "Room deleted successfully")

	})
	protectedRoomGroup.POST("/leave", func(c *gin.Context) {
		var user = controllers.Claims.GetUserFromRequest(c)
		var room = controllers.Claims.GetRoomFromRequest(c)
		if err := controllers.Room.LeaveRoom(user.Username, room.Id); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, "Room does not exist or you are not in the room")
			return
		}
		utils.SendSuccessResponse(c, http.StatusOK, nil, "Left room successfully")
	})

	roomHub := ws.NewRoomHub()
	go roomHub.Run()
	protectedRoomGroup.GET("/ws", func(c *gin.Context) {
		var room = controllers.Claims.GetRoomFromRequest(c)
		var hub = roomHub.CreateRoomHub(room.Id)
		var user = controllers.Claims.GetUserFromRequest(c)
		ws.ServeWs(hub, c.Writer, c.Request, user.Username, room.Id)
	})

	protectedRoomGroup.GET("/messages", func(c *gin.Context) {
		var room = controllers.Claims.GetRoomFromRequest(c)
		// var limit = utils.GetIntQuery(c, "limit", 10)
		// var page = utils.GetIntQuery(c, "page", 1)
		var messages = controllers.Room.GetRoomMessages(room.Id, 0, 0)
		utils.SendSuccessResponse(c, http.StatusOK, messages, "")

	})

}

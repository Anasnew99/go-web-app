package controllers

import (
	"anasnew99/server/chat_app/models"
	"anasnew99/server/chat_app/utils"

	"github.com/gin-gonic/gin"
)

func setUserClaimsInRequse(c *gin.Context, user models.User) {
	utils.SetClaimInResponse(c, "user", map[string]string{
		"username": user.Username,
		"email":    user.Email,
	})
}

func setRoomClaimsInRequest(c *gin.Context, room models.Room) {
	utils.SetClaimInResponse(c, "room", room)
}

func getUserFromRequest(c *gin.Context) models.User {
	claims, ok := utils.GetClaimFromResponse[string, models.User](c, "user")
	if !ok {
		return models.User{}
	}
	return claims
}

func getRoomFromRequest(c *gin.Context) models.Room {
	room, ok := utils.GetClaimFromResponse[string, models.Room](c, "room")
	if !ok {
		return models.Room{}
	}
	return room
}

type ClaimsController struct {
	SetUserClaimsInRequse  func(c *gin.Context, user models.User)
	SetRoomClaimsInRequest func(c *gin.Context, room models.Room)
	GetUserFromRequest     func(c *gin.Context) models.User
	GetRoomFromRequest     func(c *gin.Context) models.Room
}

var Claims = &ClaimsController{
	SetUserClaimsInRequse:  setUserClaimsInRequse,
	SetRoomClaimsInRequest: setRoomClaimsInRequest,
	GetUserFromRequest:     getUserFromRequest,
	GetRoomFromRequest:     getRoomFromRequest,
}

package controllers

import (
	"anasnew99/server/chat_app/models"
	"anasnew99/server/chat_app/utils"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Authenticate(username string, password string) (string, error) {
	fmt.Println(username, password)
	user, err := GetUser(username)
	if err != nil {
		return "", err
	}

	if user.Password != utils.GetHashedString(password) {
		return "", errors.New("wrong password")
	}
	return utils.GenerateJWT(map[string]string{
		"username": user.Username,
		"email":    user.Email,
	}, os.Getenv("JWT_SECRET"), 60*time.Minute)
}

func VerifyToken(token string) (claims map[string]string, err error) {
	return utils.VerifyJWT[string](token, os.Getenv("JWT_SECRET"))
}

func GetUserFromRequest(c *gin.Context) models.User {
	claims, ok := c.Request.Context().Value("claims").(map[string]string)
	fmt.Println(claims)
	if !ok {
		return models.User{}
	}
	return models.User{
		Username: claims["username"],
		Email:    claims["email"],
	}
}

func GetRoomFromRequest(c *gin.Context) models.Room {
	room, ok := c.Request.Context().Value("room").(models.Room)
	if !ok {
		return models.Room{}
	}
	return room
}

package controllers

import (
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
		return "", errors.New("Wrong password")
	}
	return utils.GenerateJWT(map[string]string{
		"username": user.Username,
		"email":    user.Email,
	}, os.Getenv("JWT_SECRET"), 60*time.Minute)
}

func VerifyToken(token string) (claims map[string]string, err error) {
	return utils.VerifyJWT[string](token, os.Getenv("JWT_SECRET"))
}

func GetUserFromRequest(c *gin.Context) User {
	claims, ok := c.Request.Context().Value("claims").(map[string]string)
	fmt.Println(claims)
	if !ok {
		return User{}
	}
	return User{
		Username: claims["username"],
		Email:    claims["email"],
	}
}

package controllers

import (
	"anasnew99/server/chat_app/utils"
	"errors"
	"fmt"
	"os"
	"time"
)

type AuthController struct {
	Authenticate func(username string, password string) (string, error)
	VerifyToken  func(token string) (map[string]string, error)
}

var Auth = &AuthController{
	Authenticate: func(username string, password string) (string, error) {
		fmt.Println(username, password)
		user, err := User.GetUser(username)
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
	},
	VerifyToken: func(token string) (claims map[string]string, err error) {
		return utils.VerifyJWT[string](token, os.Getenv("JWT_SECRET"))
	},
}

package routers

import (
	"anasnew99/server/chat_app/constants"
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/models"
	"anasnew99/server/chat_app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAuthRouter(r *gin.RouterGroup) {
	r.POST("/login", func(c *gin.Context) {
		var credentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&credentials); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, fmt.Sprintf("%v", err))
			return
		}

		token, err := controllers.Auth.Authenticate(credentials.Username, credentials.Password)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, constants.BAD_CREDENTIALS, "Username or password is incorrect")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, gin.H{
			"token": token,
		}, "")

	})

	r.POST("/register", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, fmt.Sprintf("%v", err))
			return
		}

		if _, err := controllers.User.AddUser(user); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, "Username already exists")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, nil, "User created successfully")

	})
}

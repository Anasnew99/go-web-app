package routers

import (
	"anasnew99/server/chat_app/constants"
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRouter(r *gin.RouterGroup) {
	r.GET("/profile", func(c *gin.Context) {
		var user = controllers.Claims.GetUserFromRequest(c)
		user, err := controllers.User.GetUser(user.Username)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusNotFound, constants.NOT_FOUND, "User does not exist")
			return
		}
		user.Password = ""
		utils.SendSuccessResponse(c, http.StatusOK, user, "")
	})

	r.PATCH("/password", func(c *gin.Context) {
		var user = controllers.Claims.GetUserFromRequest(c)

		var password struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}

		if err := c.ShouldBindJSON(&password); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, err.Error())
			return
		}

		if err := controllers.User.ChangeUserPassword(user.Username, password.OldPassword, password.NewPassword); err != nil {
			utils.SendErrorResponse(c, http.StatusBadRequest, constants.INVALID_DATA, "Old password is incorrect")
			return
		}

		utils.SendSuccessResponse(c, http.StatusOK, nil, "Password changed successfully")

	})
}

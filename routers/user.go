package routers

import (
	"anasnew99/server/chat_app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRouter(r *gin.RouterGroup) {
	r.GET("/profile", func(c *gin.Context) {
		var user = controllers.GetUserFromRequest(c)

		user, err := controllers.GetUser(user.Username)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		user.Password = ""
		c.JSON(http.StatusOK, gin.H{"user": user})
	})

	r.PATCH("/password", func(c *gin.Context) {
		var user = controllers.GetUserFromRequest(c)

		var password struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}

		if err := c.ShouldBindJSON(&password); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := controllers.ChangeUserPassword(user.Username, password.OldPassword, password.NewPassword); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})

	})
}

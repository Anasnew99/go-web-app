package routers

import (
	"anasnew99/server/chat_app/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRouter(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		var user = controllers.GetUserFromRequest(c)

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %s, will see user results here soon", user.Username),
		})
	})
}

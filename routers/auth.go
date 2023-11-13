package routers

import (
	"anasnew99/server/chat_app/controllers"
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
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Bad request: %v", err),
			})
			return
		}

		token, err := controllers.Authenticate(credentials.Username, credentials.Password)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Unauthorized: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	})
}

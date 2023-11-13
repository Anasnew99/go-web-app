package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addUserRouter(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World, You will see user result here!",
		})
	})
}

package middlewares

import (
	"anasnew99/server/chat_app/controllers"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdminRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authorization = c.GetHeader("Authorization")
		if len(authorization) <= len("Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Unauthorized: %v", "Invalid token"),
			})
			return
		}
		var bearerToken = authorization[len("Bearer "):]

		claims, err := controllers.VerifyToken(bearerToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Unauthorized: %v", err),
			})
			return
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "claims", claims))

		c.Next()
	}
}

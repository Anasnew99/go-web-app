package middlewares

import (
	"anasnew99/server/chat_app/controllers"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthAdminRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authorization = strings.Trim(c.GetHeader("Authorization"), " ")
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

func AuthRoomRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user = controllers.GetUserFromRequest(ctx)
		var roomID = ctx.Param("id")
		var room, isJoined = controllers.IsUserJoinedInTheRoom(user.Username, roomID)
		if !isJoined {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Unauthorized: %v", "You are not joined in this room"),
			})
			return
		}
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "room", room))
		ctx.Next()

	}
}

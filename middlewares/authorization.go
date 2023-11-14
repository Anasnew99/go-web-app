package middlewares

import (
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/models"
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

		claims, err := controllers.Auth.VerifyToken(bearerToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Unauthorized: %v", err),
			})
			return
		}
		controllers.Claims.SetUserClaimsInRequse(c, models.User{
			Username: claims["username"],
			Email:    claims["email"],
		})
		c.Next()
	}
}

func AuthRoomRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user = controllers.Claims.GetUserFromRequest(ctx)
		var roomID = ctx.Param("id")
		var room, isJoined = controllers.Room.IsUserJoinedInTheRoom(user.Username, roomID)
		if !isJoined {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprintf("Unauthorized: %v", "You are not joined in this room"),
			})
			return
		}
		controllers.Claims.SetRoomClaimsInRequest(ctx, room)
		ctx.Next()

	}
}

package middlewares

import (
	"anasnew99/server/chat_app/constants"
	"anasnew99/server/chat_app/controllers"
	"anasnew99/server/chat_app/models"
	"anasnew99/server/chat_app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthAdminRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authorization = strings.Trim(c.GetHeader("Authorization"), " ")

		if len(authorization) <= len("Bearer ") {
			// if it is websocket connection then take token from query
			if strings.Contains(c.FullPath(), "/ws") {
				if c.Query("token") == "" {
					utils.SendErrorResponse(c, http.StatusUnauthorized, constants.TOKEN_EXPIRED, "Please login again. Token is invalid")
					return
				}
				authorization = "Bearer " + c.Query("token")

			} else {
				utils.SendErrorResponse(c, http.StatusUnauthorized, constants.TOKEN_EXPIRED, "Please login again. Token is invalid")
				return

			}

		}
		var bearerToken = authorization[len("Bearer "):]

		claims, err := controllers.Auth.VerifyToken(bearerToken)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, constants.TOKEN_EXPIRED, "Please login again. Token expired")
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
			utils.SendErrorResponse(ctx, http.StatusForbidden, constants.UNAUTHORIZED, "You are not joined in the room")
			return
		}
		controllers.Claims.SetRoomClaimsInRequest(ctx, room)
		ctx.Next()

	}
}

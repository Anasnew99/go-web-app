package routers

import (
	"anasnew99/server/chat_app/middlewares"

	"github.com/gin-gonic/gin"
)

func addProtectedRoutes(r *gin.RouterGroup) {
	r.Use(middlewares.AuthAdminRequest())
	addUserRouter(r.Group("/user"))
}

func AddRouters(r *gin.Engine) {
	addProtectedRoutes(r.Group("/admin"))
}

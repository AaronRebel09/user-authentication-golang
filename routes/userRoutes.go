package routes

import (
	controller "github.com/AaronRebel09/user-authentication-golang/controllers"
	"github.com/AaronRebel09/user-authentication-golang/middleware"
	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/signin", controller.Login())
	//Se agrega middleware Authntication solo a este endpoint
	incomingRoutes.GET("/users/info", middleware.Authentication(), controller.SimpleEndpoint())
	incomingRoutes.POST("/users/upload", middleware.Authentication(), controller.UploadImages())
}

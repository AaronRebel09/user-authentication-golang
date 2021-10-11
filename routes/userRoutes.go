package routes

import (
	controller "github.com/AaronRebel09/user-authentication-golang/controllers"
	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/signin",controller.Login())
}
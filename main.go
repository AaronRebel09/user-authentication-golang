package main

import (
	"os"

	middleware "github.com/AaronRebel09/user-authentication-golang/middleware"
	routes "github.com/AaronRebel09/user-authentication-golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	//Se agrega Cross Origin Resource Sharing
	router.Use(middleware.CORSMiddleware())
	routes.UserRoutes(router)
	// API-1
	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api v1"})
	})

	// API-2
	router.GET("/api/v2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api v2"})
	})
	router.Run(":" + port)

}

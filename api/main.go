package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/api/controllers"
	"github.com/lucapierini/api/initializers"
	"github.com/lucapierini/api/middleware"
)

func init(){
	initializaers.LoadEnvVariables()
	initializaers.ConnectToDb()
	initializaers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("pong %s", "🏓"),
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}

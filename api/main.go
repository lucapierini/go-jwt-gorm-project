package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucapierini/api/initializaers"
)

func init(){
	initializaers.LoadEnvVariables()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("pong %s", "🏓"),
		})
	})
	r.Run()
}

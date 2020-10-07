package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		message := fmt.Sprintf("Hello %s", name)
		c.JSON(200, gin.H{"message": message})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}

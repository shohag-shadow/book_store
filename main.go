package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from gin"})
	})
	r.GET("/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"time": time.Now().String()})
	})
	r.Run(":8080")
}

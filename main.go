package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(Options)

	r.GET("/user/:name/test", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name|test": name,
		})
	})
	r.GET("/user/:name/passwd/:passwd", func(c *gin.Context) {
		name := c.Param("name")
		passwd := c.Param("passwd")
		c.JSON(http.StatusOK, gin.H{
			"name|test":   name,
			"name|passwd": passwd,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

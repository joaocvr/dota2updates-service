package main

import (
	"fmt"

	"dota2updates-service/endpoint"

	"github.com/gin-gonic/gin"
)

func main() {
	a := endpoint.Api{}

	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/v1/news", a.GetNews)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Print(err.Error())
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

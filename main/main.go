package main

import (
	"fmt"

	"dota2updates-service/endpoint"

	"github.com/gin-gonic/gin"
)

func main() {
	a := endpoint.Api{}

	router := gin.Default()
	router.GET("/news", a.GetNews)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Print(err.Error())
	}
}

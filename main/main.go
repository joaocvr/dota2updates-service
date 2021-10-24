package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joaocvr/endpoint"
)

func main() {
	a := endpoint.API{}

	router := gin.Default()
	router.GET("/news", a.getNews)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Print(err.Error())
	}
}

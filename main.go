package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/")

	address := "localhost:8000"
	err := router.Run(address)
	if err != nil {
		fmt.Println("Failed to run server at address " + address)
		return
	}
}

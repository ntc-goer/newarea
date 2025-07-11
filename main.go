package main

import (
	"github.com/gin-gonic/gin"
	"newarea/handler"
)

func main() {
	router := gin.Default()

	healthHandler := handler.NewHealthHandler()

	router.GET("/health", healthHandler.Health)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

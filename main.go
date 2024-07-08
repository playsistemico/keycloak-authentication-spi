package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main.go/handlers"
	"main.go/services"
)

func main() {
	router := gin.Default()

	authService := &services.MockAuthService{}

	loginHandler := handlers.NewLoginHandler(authService)
	app2AppAuthHandler := handlers.NewApp2AppAuthHandler(authService)

	bankGroup := router.Group("/backend-bank")

	{
		bankGroup.POST("/login", loginHandler.HandleLogin)
		bankGroup.POST("/app2app/auth", app2AppAuthHandler.HandleApp2AppAuth)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run("127.0.0.1:9090")
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main.go/services"
)

type App2AppAuthHandler struct {
	AuthService services.AuthService
}

func NewApp2AppAuthHandler(authService services.AuthService) *App2AppAuthHandler {
	return &App2AppAuthHandler{
		AuthService: authService,
	}
}

func (h *App2AppAuthHandler) HandleApp2AppAuth(c *gin.Context) {
	var authData struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&authData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.AuthService.ValidateToken(authData.Token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Everything is OK"})
}

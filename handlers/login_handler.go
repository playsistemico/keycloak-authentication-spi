package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"main.go/services"
)

type LoginHandler struct {
	AuthService services.AuthService
}

func NewLoginHandler(authService services.AuthService) *LoginHandler {
	return &LoginHandler{
		AuthService: authService,
	}
}

func (h *LoginHandler) HandleLogin(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.Authenticate(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/defines"
	services "backend/internal/service"
)

type LoginController interface {
	Login(ctx *gin.Context)
	ValidateSession(ctx *gin.Context)
}
type loginController struct {
	srv services.LoginService
}

func NewLoginController(srv services.LoginService) LoginController {
	return &loginController{
		srv: srv,
	}
}

func (ctrl *loginController) Login(c *gin.Context) {
	var loginData defines.LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.srv.Login(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctrl *loginController) ValidateSession(c *gin.Context) {
	var authData defines.AuthData

	if err := c.ShouldBindJSON(&authData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ctrl.srv.ValidateSession(authData.Token)
	if !resp || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Everything is OK"})

}

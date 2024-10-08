package controller

import (
	"backend/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"

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
	var loginData domain.LoginRequestBody

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := ctrl.srv.Create(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	resp := domain.LoginResponseBody{Session: session}
	c.JSON(http.StatusOK, resp)
}

func (ctrl *loginController) ValidateSession(c *gin.Context) {

	session := c.Query("session")
	if session == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session parameter is required"})
		return
	}

	resp, err := ctrl.srv.GetUser(session)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"username": resp.Name})

}

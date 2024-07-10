package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/controller"
	"backend/internal/defines"
	services "backend/internal/service"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {

	// Services init
	loginSvc := &services.MockLoginService{}

	// Controllers init
	loginCtrl := controller.NewLoginController(loginSvc)

	// Endpoints
	r.GET(defines.EndpointPing, healthCheck)
	r.POST(defines.EndpointLogin, loginCtrl.Login)
	r.POST(defines.EndpointApp2app, loginCtrl.ValidateSession)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

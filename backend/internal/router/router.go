package router

import (
	"backend/internal/service"
	"net/http"

	"backend/internal/controller"
	"backend/internal/defines"
	"backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// Repository init
	usersRepo := repository.NewUsersRepository()

	// Services init
	loginSvc := service.NewLoginService(usersRepo)

	// Controllers init
	loginCtrl := controller.NewLoginController(loginSvc)

	// Endpoints
	r.GET(defines.EndpointPing, healthCheck)
	r.POST(defines.EndpointLogin, loginCtrl.Login)
	r.GET(defines.EndpointUser, loginCtrl.ValidateSession)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

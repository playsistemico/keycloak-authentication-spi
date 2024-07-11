package router

import (
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"backend/db"
	"backend/internal/controller"
	"backend/internal/defines"
	"backend/internal/repository"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	database := db.InitDB()
	sqlxDB := sqlx.NewDb(database, "sqlite3")

	// Repository init
	usersRepo := repository.NewUsersRepository(sqlxDB)

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

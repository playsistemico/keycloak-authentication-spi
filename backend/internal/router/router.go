package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"backend/db"
	"backend/internal/controller"
	"backend/internal/defines"
	"backend/internal/repository"
	services "backend/internal/service"
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
	loginSvc := services.NewLoginService(usersRepo)

	// Controllers init
	loginCtrl := controller.NewLoginController(loginSvc)

	// Endpoints
	r.GET(defines.EndpointPing, healthCheck)
	r.POST(defines.EndpointLogin, loginCtrl.Login)
	r.GET(defines.EndpointApp2app, loginCtrl.ValidateSession)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

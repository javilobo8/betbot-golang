package controllers

import (
	"betbot/db"
	"betbot/services"

	"github.com/gin-gonic/gin"
)

var DBClient = db.ConnectDB()

var UserRepository = services.NewUserRepository(*DBClient)

func BindControllers(app *gin.Engine) {
	BindLoginController(app)
}

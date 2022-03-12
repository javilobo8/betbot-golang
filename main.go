package main

import (
	"betbot/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	app := gin.New()
	app.SetTrustedProxies(nil)
	controllers.BindControllers(app)
	return app
}

func initApp(app *gin.Engine) {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Init")

	app := setUpRouter()

	initApp(app)
}

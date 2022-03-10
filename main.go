package main

import (
	"betbot/constants"
	"betbot/controllers"
	"betbot/util"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/gin-gonic/gin"
)

var botUser = os.Getenv("TWITCH_CHAT_USER")
var botPass = os.Getenv("TWITCH_CHAT_PASS")

var channels []string = []string{
	"javilobo8",
}

func onPrivateMessage(message twitch.PrivateMessage) {
	if message.Message[0] == constants.CommandChar {
		messages := util.GetCommandMessages(message.Message)

		switch messages[0] {
		case constants.BetCommand:
			{
				log.Println(message.Message, messages)
				break
			}
		}
	}
}

func onConnect() {
	log.Println("Bot connected")
}

func setUpRouter() *gin.Engine {
	app := gin.New()
	app.SetTrustedProxies(nil)
	app.GET("/ping", controllers.PingHandlerGET)
	return app
}

func initApp(app *gin.Engine) {
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func initBot(client *twitch.Client) {
	log.Println("pre-twitch")
	client.Join(channels...)
	client.OnConnect(onConnect)
	client.OnPrivateMessage(onPrivateMessage)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Init")
	var client *twitch.Client
	if botUser != "" && botPass != "" {
		client = twitch.NewClient(botUser, botPass)
	} else {
		client = twitch.NewAnonymousClient()
	}

	app := setUpRouter()

	go initBot(client)
	initApp(app)
}

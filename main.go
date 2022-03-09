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
		messages := util.GetMessages(message.Message)

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

func initApp(app *gin.Engine) {
	log.Println("pre-app")
	app.SetTrustedProxies(nil)

	app.GET("/ping", controllers.PingHandlerGET)

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
	client := twitch.NewClient(botUser, botPass)
	app := gin.New()

	go initBot(client)
	initApp(app)
}

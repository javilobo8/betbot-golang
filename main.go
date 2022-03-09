package main

import (
	"betbot/constants"
	"betbot/util"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/gin-gonic/gin"
)

var botUser = os.Getenv("TWITCH_CHAT_USER")
var botPass = os.Getenv("TWITCH_CHAT_PASS")

var client *twitch.Client
var app *gin.Engine

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

func initApp() {
	log.Println("pre-app")
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func initBot() {
	log.Println("pre-twitch")
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Init")
	gin.SetMode(gin.DebugMode)

	client = twitch.NewClient(botUser, botPass)
	app = gin.New()

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	client.Join(channels...)
	client.OnConnect(onConnect)
	client.OnPrivateMessage(onPrivateMessage)

	go initBot()
	initApp()
}

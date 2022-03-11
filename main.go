package main

import (
	"betbot/constants"
	"betbot/controllers"
	"betbot/models"
	"betbot/util"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func testMongo(mongoClient *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database := mongoClient.Database("go-test")
	usersCollection := database.Collection("users")

	insertResult, err := usersCollection.InsertOne(ctx, models.User{
		TwitchId: 12345,
		UserName: "test-username",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	log.Println("Init")
	mongoClient := ConnectDB()
	var client *twitch.Client
	if botUser != "" && botPass != "" {
		client = twitch.NewClient(botUser, botPass)
	} else {
		client = twitch.NewAnonymousClient()
	}

	app := setUpRouter()

	testMongo(mongoClient)

	go initBot(client)
	initApp(app)
}

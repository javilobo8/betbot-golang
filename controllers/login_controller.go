package controllers

import (
	"betbot/models"
	"betbot/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

func BindLoginController(app *gin.Engine) {
	app.POST("/login", Login)
	app.POST("/login/create-user", CreateUser)
}

func Login(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "'username' is required"})
		return
	}
	password := ctx.Query("password")
	if password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "'password' is required"})
		return
	}

	log.Println("body", username, "password", password)

	user := UserRepository.FindByUsername(username)
	if user == (models.User{}) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !util.CheckPasswordHash(password, user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong password"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var body CreateUserRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.Validate(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := UserRepository.FindByUsername(body.Username)

	if user != (models.User{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	user = UserRepository.Create(body.Username, body.Password)

	ctx.JSON(http.StatusOK, user)
}

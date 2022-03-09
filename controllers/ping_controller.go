package controllers

import "github.com/gin-gonic/gin"

func PingHandlerGET(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}

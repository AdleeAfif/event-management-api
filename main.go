package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", getDefault)

	server.Run(":8080")
}

func getDefault(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World"}) //(200, "Hello World")
}

package main

import (
	"net/http"
	"project/event-management-api/db"
	"project/event-management-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", getDefault)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getDefault(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World"}) //(200, "Hello World")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"event": event})
}

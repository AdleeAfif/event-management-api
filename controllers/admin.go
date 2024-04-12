package controllers

import (
	"net/http"
	"project/event-management-api/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	context.JSON(http.StatusOK, users)
}

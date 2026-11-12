package controller

import (
	"net/http"

	"example.exe/event-app/config"
	"example.exe/event-app/models"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	userId, _ := context.Get("userId")
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event.UserId = userId.(int)
	config.DB.Create(&event)

	context.JSON(http.StatusCreated, gin.H{
		"message": "data berhasil dibuat",
		"event":   event,
	})
}

func GetEvents(context *gin.Context) {
	var events []models.Event

	config.DB.Find(&events)
	context.JSON(http.StatusOK, gin.H{
		"message": "data berhasil dibuat",
		"event":   events,
	})
	return
}

func GetEventbyId(context *gin.Context) {
	var event models.Event
	paramsId := context.Param("id")
	var eventData = config.DB.First(&event, paramsId).Error

	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event tidak ditemukan",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "data tampil detail event",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {

	userId, _ := context.Get("userId")

	var event models.Event
	paramsId := context.Param("id")
	var eventData = config.DB.First(&event, paramsId).Error

	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event tidak ditemukan",
		})
		return
	}

	if event.UserId != userId.(int) {
		context.JSON(http.StatusForbidden, gin.H{
			"error": "Anda tidak bisa mengupdate Event user lain",
		})
		return
	}

	var input models.Event

	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Model(&event).Updates(input)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil di update",
		"event":   event,
	})
	
}

func DeleteEvent(context *gin.Context) {

	userId, _ := context.Get("userId")

	var event models.Event
	paramsId := context.Param("id")
	var eventData = config.DB.First(&event, paramsId).Error

	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event tidak ditemukan",
		})
		return
	}

	if event.UserId != userId.(int) {
		context.JSON(http.StatusForbidden, gin.H{
			"error": "Anda tidak bisa menghapus Event orang lain",
		})
		return
	}

	config.DB.Unscoped().Delete(&event)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil di delete",
		"event":   event,
	})
}

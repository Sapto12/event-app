package main

import (
	"net/http"
	"time"

	"example.exe/event-app/models"
	"github.com/gin-gonic/gin"
)



func main(){
	server := gin.Default()


	//route
	api:= server.Group("/api")
	{

		api.POST("/events", createEvents)
		api.GET("/events", getEvents)
	}

	server.Run(":8080")



}

//untuk handler

func getEvents(context *gin.Context){
	events:= models.GetAllEvents()

	context.JSON(http.StatusOK, events)
}


func createEvents(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)


	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"Message" : "cound not parse req data",
			"err" : err.Error(),
		})
		return
	}
	//dummy
	tm:= time.Now()
	event.Id = 1 
	event.UserId = 1
	event.DateTime = tm

	event.Save();

	context.JSON(http.StatusCreated, 
	gin.H{
		"Message" : "Create Event",
		"event" : event,
	})
}
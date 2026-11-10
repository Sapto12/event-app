package main

import (
	"log"
	"example.exe/event-app/config"
	"example.exe/event-app/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main(){

	err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file")
}

	config.ConnectDB()
	server := gin.Default()


	//route
	api:= server.Group("/api")
	{

		api.POST("/events", controller.CreateEvent)
		api.GET("/events", controller.GetEvents)
		api.GET("/events/:id",controller.GetEventbyId )
	}

	server.Run(":8080")



}


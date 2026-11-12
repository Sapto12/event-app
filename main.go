package main

import (
	"log"

	"example.exe/event-app/config"
	"example.exe/event-app/controller"
	"example.exe/event-app/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	server := gin.Default()

	//route
	api := server.Group("/api")
	{
		//crud route
		
		api.GET("/events", controller.GetEvents)
		api.GET("/events/:id", controller.GetEventbyId)
		

		//login register route
		api.POST("/auth/register", controller.RegisterUser)
		api.POST("/auth/login", controller.LoginUser)

		//middleware

		protected := api.Group("/")
		protected.Use(middleware.RequiredAuth())
		{
			protected.GET("/auth/me", controller.GetCurrentUser)
			protected.POST("/events", controller.CreateEvent)
			protected.PUT("/events/:id", controller.UpdateEvent)
			protected.DELETE("/events/:id", controller.DeleteEvent)
		}
	}

	server.Run(":8080")

}

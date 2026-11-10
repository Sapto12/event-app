package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func main(){
	server := gin.Default()


	//route 

	server.GET("/events", getEvent)
	server.GET("/test", getTest)

	server.Run(":8080")



}

//untuk handler

func getEvent(context *gin.Context){

	context.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
	
}


func getTest(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
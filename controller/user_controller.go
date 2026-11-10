package controller

import (
	"net/http"

	"example.exe/event-app/config"
	"example.exe/event-app/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthInputRegister struct {
	Name     string `json : "name" binding:"required"`
	Email    string `json : "email" binding:"required, email"`
	Password string `json : "password" binding:"required, min=6"`
}
type AuthInputLogin struct {
	Email    string `json : "email" binding:"required, email"`
	Password string `json : "password" binding:"required"`
}

func RegisterUser(context * gin.Context){
	var input AuthInputRegister

	//validation

	err:= context.ShouldBindJSON(&input);
	if err != nil{
		context.JSON(http.StatusBadRequest,
		gin.H{
			"error": err.Error(),
		})
		return
	}
	//hash password

	hashedPassword, errhash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if errhash != nil{
		context.JSON(http.StatusBadRequest,
		gin.H{
			"error": "Gagal enkripisi password",

		})
		return	
	}


	//simpan ke database

	user:=models.User{
		Name: input.Name,
		Email: input.Email,
		Password: string(hashedPassword),
	}


	userCreated := config.DB.Create(&user).Error

	if userCreated != nil{
		context.JSON(http.StatusBadRequest,
			gin.H{
				"error" : "Email sudah terdaftar",
			})	
	}

	context.JSON(http.StatusCreated, gin.H{
		"message" : "Berhasil register",
		"user" : gin.H{
			"id" : user.ID,
			"name" : user.Name,
			"email" : user.Email,
			"events" : user.Events,
		},
	})
}
package controller

import (
	"net/http"
	"os"
	"time"

	"example.exe/event-app/config"
	"example.exe/event-app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthInputRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthInputLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(context *gin.Context) {
	var input AuthInputRegister

	// Validasi input
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Hash password
	hashedPassword, errhash := bcrypt.GenerateFromPassword(
		[]byte(input.Password),
		bcrypt.DefaultCost,
	)

	if errhash != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengenkripsi password",
		})
		return
	}

	// Simpan ke database
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	userCreated := config.DB.Create(&user).Error

	if userCreated != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal membuat akun. Email mungkin sudah terdaftar.",
		})
		return
	}

	// Respons berhasil
	context.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil register",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func LoginUser(context *gin.Context) {
	var input AuthInputLogin

	// Validasi input
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//cek terdaftar atau tidak
	var user models.User
	userData := config.DB.Where("email = ?", input.Email).First(&user).Error
	if userData != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "email belum terdaftar ",
		})
		return
	}
	errMatchPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if errMatchPassword != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "password salah",
		})
		return
	}
	//buat token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Gagal Membuat token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "login berhasil",
		"token":   tokenString,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})

}

func GetCurrentUser(context*gin.Context){
	//ambil user id
	
	userId, exists := context.Get("userId")
		if !exists{
			context.JSON(http.StatusInternalServerError, gin.H{
				"error" : "tidak authentication",
			})
			return
		}
		var   user models.User

		userData := config.DB.Select("id", "name", "email").First(&user, userId).Error
		if userData != nil{
			context.JSON(http.StatusNotFound, gin.H{
				"error" : "user tidak ditemukan",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"user" : gin.H{
				"id" : user.ID,
				"name" : user.Name,
				"email" : user.Email,
				"events" : user.Events,
			},
		})
	}

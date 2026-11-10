package config

import (
	"log"
	"os"
	
	"gorm.io/driver/postgres"
	"example.exe/event-app/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := os.Getenv("DATABASE_URI")
	if dsn == ""{
		log.Fatal("Environment variable belum diisi")

	}

	database, err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})


	if err != nil {
		log.Fatal("Gagal terkoneksi ke database", err)


	}

	err = database.AutoMigrate(&models.Event{}, models.User{})
	if err != nil{
		log.Fatal("Gagal melakukan migration database", err)
	}

	DB = database
	log.Println("berhasil terkoneksi ke database")
}
package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //this variable hold he database connection

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Get("DB_HOST"),
		Get("DB_USER"),
		Get("DB_PASSWORD"),
		Get("DB_NAME"),
		Get("DB_PORT"),
	)
	//this line is opening the connection to postgresql through gorm
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = db
	log.Println("Database connected successfully")
}

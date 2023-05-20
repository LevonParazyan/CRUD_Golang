package initializers

import (
	"go-crud/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb()  {
	// Retrieve environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Database connection string
	dsn := fmt.Sprintf("%s@tcp(%s:%s)/?parseTime=true", dbUser, dbHost, dbPort)

	// Connect to MySQL server
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create the 'users' database if it doesn't exist
	err = DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	// Connect to the 'users' database
	dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbHost, dbPort, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Automatically migrate the User model to the database
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate User model:", err)
	}

	fmt.Println("Connected to MySQL!")

}
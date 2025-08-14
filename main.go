package main

import (
	"log"

	"github.com/hale-services/config"
	models "github.com/hale-services/models/users"
	"github.com/hale-services/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.UserFitnessProfile{})

	r := router.SetUpRouter()
	// Set up the routes
	r.Run(":8080")
}

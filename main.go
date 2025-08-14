package main

import (
	"github.com/axhutoxh/go-starter/config"
	models "github.com/axhutoxh/go-starter/models/users"
	"github.com/axhutoxh/go-starter/router"
)

func main() {
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{}, &models.UserFitnessProfile{})

	r := router.SetUpRouter()
	// Set up the routes
	r.Run(":8080")
}

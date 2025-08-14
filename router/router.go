package router

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/hale-services/handlers/users"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/user/:id", handlers.GetUser)
		api.POST("/users", handlers.CreateUser)
		api.PUT("/users/:id", handlers.UpdateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)
	}

	return r
}

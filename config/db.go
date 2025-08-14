package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GORM is a popular Object Relational Mapping (ORM) library for the Go programming language.
// It provides an elegant and developer-friendly way to interact with databases by allowing you to work directly with Go structs
// instead of writing raw SQL queries.
var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=nopassword dbname=demo port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Get the generic database object sql.DB to configure connection pool
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatal("failed to get database object:", err)
	}

	// Optimize connection pool settings for performance
	sqlDB.SetMaxOpenConns(50)           // Increase for high concurrency
	sqlDB.SetMaxIdleConns(10)           // Set idle connections for reuse
	sqlDB.SetConnMaxLifetime(time.Hour) // Limit connection lifetime to avoid stale connections

	DB = database
}

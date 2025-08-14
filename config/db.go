package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GORM is a popular Object Relational Mapping (ORM) library for the Go programming language.
// It provides an elegant and developer-friendly way to interact with databases by allowing you to work directly with Go structs
// instead of writing raw SQL queries.
var DB *gorm.DB

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

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

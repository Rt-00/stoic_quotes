package main

import (
	"stoic_quotes/db"
	"stoic_quotes/handlers"
	"stoic_quotes/repository"
	"stoic_quotes/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connecting to the database
	database := db.Connect()
	defer database.Close()

	// Import CSV file
	repo := repository.QuoteRepository{DB: database}
	service.ImportCSV(&repo, "./stoic_quotes_full.csv")

	// Create Gin Server Engine
	r := gin.Default()

	// CORS config
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}

	r.Use(cors.New(config))

	// Routes
	r.GET("/quotes", handlers.ListQuotes(database))
	r.GET("/quotes/random", handlers.RandomQuote(database))

	// Start server
	r.Run(":3000")
}

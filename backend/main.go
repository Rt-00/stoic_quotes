package main

import (
	"stoic_quotes/db"
	"stoic_quotes/repository"
	"stoic_quotes/service"
)

func main() {
	// Connecting to the database
	database := db.Connect()
	defer database.Close()

	// Import CSV file
	repo := repository.QuoteRepository{DB: database}
	service.ImportCSV(&repo, "./stoic_quotes_full.csv")
}

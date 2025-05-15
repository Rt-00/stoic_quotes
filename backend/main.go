package main

import "stoic_quotes/db"

func main() {
	// Connecting to the database
	database := db.Connect()
	defer database.Close()
}

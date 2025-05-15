package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./quotes.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS quotes (
		id     INTEGER PRIMARY KEY AUTOINCREMENT,
		quote TEXT    UNIQUE,
		author TEXT,
		book   TEXT
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("error creating quotes table: %s", err)
	}

	createQuotesUniqueIndex := `
	CREATE UNIQUE INDEX IF NOT EXISTS uidx_quotes_quote ON quotes(quote);
	`
	_, err = db.Exec(createQuotesUniqueIndex)
	if err != nil {
		log.Fatalf("error creating unique index for quotes: %s", err)
	}

	return db
}

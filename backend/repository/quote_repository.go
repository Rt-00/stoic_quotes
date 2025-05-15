package repository

import (
	"database/sql"
	"fmt"
	"log"
	"stoic_quotes/models"
	"strings"
)

type QuoteRepository struct {
	DB *sql.DB
}

func (r *QuoteRepository) InsertBatch(quotes []models.Quote) {
	if len(quotes) == 0 {
		return
	}

	var builder strings.Builder
	builder.WriteString("INSERT OR IGNORE INTO quotes (quote, author, book) VALUES ")

	var params []any
	for i, q := range quotes {
		if i > 0 {
			builder.WriteString(",")
		}

		builder.WriteString("(?, ?, ?)")
		params = append(params, q.Quote, q.Author, q.Book)
	}

	_, err := r.DB.Exec(builder.String(), params...)
	if err != nil {
		log.Printf("Batch insert failed: %v", err)
	} else {
		fmt.Printf("Inserted batch of %d quotes\n", len(quotes))
	}
}

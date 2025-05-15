package service

import (
	"encoding/csv"
	"log"
	"os"
	"stoic_quotes/models"
	"stoic_quotes/repository"
)

const BATCH_SIZE = 500

func ImportCSV(repo *repository.QuoteRepository, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s: %s", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read file: %s: %s", filePath, err)
	}

	var batch []models.Quote
	for i, row := range rows {
		if i == 0 {
			continue
		}

		q := models.Quote{
			Quote:  row[0],
			Author: row[1],
			Book:   row[2],
		}

		batch = append(batch, q)

		if len(batch) >= BATCH_SIZE {
			repo.InsertBatch(batch)
			batch = nil
		}
	}

	if len(batch) > 0 {
		repo.InsertBatch(batch)
	}
}

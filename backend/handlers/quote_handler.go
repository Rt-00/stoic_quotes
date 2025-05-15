package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"stoic_quotes/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListQuotes(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.Query("author")
		book := c.Query("book")
		query := c.Query("q")
		page := c.DefaultQuery("page", "1")
		perPage := c.DefaultQuery("per_page", "10")

		var where []string
		var args []any

		if author != "" {
			where = append(where, "LOWER(author) LIKE ?")
			args = append(args, "%"+strings.ToLower(author)+"%")
		}
		if book != "" {
			where = append(where, "LOWER(book) LIKE ?")
			args = append(args, "%"+strings.ToLower(book)+"%")
		}
		if query != "" {
			where = append(where, "LOWER(quote) LIKE ?")
			args = append(args, "%"+strings.ToLower(query)+"%")
		}

		whereSQL := ""
		if len(where) > 0 {
			whereSQL = "WHERE " + strings.Join(where, " AND ")
		}

		// Count total results
		var totalResults int
		countQuery := fmt.Sprintf("SELECT COUNT(*) FROM quotes %s", whereSQL)
		err := db.QueryRow(countQuery, args...).Scan(&totalResults)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Paginação
		var pageInt, perPageInt int
		fmt.Sscanf(page, "%d", &pageInt)
		fmt.Sscanf(perPage, "%d", &perPageInt)
		offset := (pageInt - 1) * perPageInt

		sqlQuery := fmt.Sprintf(`
			SELECT
			  id,
				quote,
				author,
				book
			FROM
			  quotes
			%s
			LIMIT ?
			OFFSET ?
		`, whereSQL)

		argsWithLimit := append(args, perPageInt, offset)

		rows, err := db.Query(sqlQuery, argsWithLimit...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		quotes := []models.Quote{}
		for rows.Next() {
			var q models.Quote
			rows.Scan(&q.ID, &q.Quote, &q.Author, &q.Book)
			quotes = append(quotes, q)
		}

		totalPages := (totalResults + perPageInt - 1) / perPageInt // Round Up

		c.JSON(http.StatusOK, gin.H{
			"data":         quotes,
			"total":        totalResults,
			"total_pages":  totalPages,
			"current_page": pageInt,
			"per_page":     perPageInt,
		})
	}
}

func RandomQuote(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		row := db.QueryRow(`
		SELECT
		  id,
			quote,
			author,
			book
		FROM
			quotes
		ORDER BY
			RANDOM()
		LIMIT
			1
		`)

		var q models.Quote
		err := row.Scan(&q.ID, &q.Quote, &q.Author, &q.Book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, q)
	}
}

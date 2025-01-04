// package main

// import (
// 	"crypto/rand"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// const (
// 	rowCount = 10000000 // Adjust for 1K, 1M, or 10M rows
// )

// func main() {
// 	// Connect to the database
// 	db, err := sql.Open("mysql", "root:Saibaba@2024@tcp(127.0.0.1:3306)/url_shortner")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	defer db.Close()

// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("Database connection error: %v", err)
// 	}
// 	log.Println("Database connection successful!")

// 	// Prepare the insert statement
// 	stmt, err := db.Prepare("INSERT IGNORE INTO url_shortner (original_url, short_code) VALUES (?, ?)")
// 	if err != nil {
// 		log.Fatalf("Failed to prepare statement: %v", err)
// 	}
// 	defer stmt.Close()

// 	start := time.Now()
// 	log.Printf("Starting insertion of %d rows...\n", rowCount)

// 	for i := 1; i <= rowCount; i++ {
// 		originalURL := fmt.Sprintf("https://example.com/long-url-%d", i)
// 		shortCode := generateUniqueShortCode()

// 		_, err := stmt.Exec(originalURL, shortCode)
// 		if err != nil {
// 			log.Printf("Error inserting row %d: %v\n", i, err)
// 		}

// 		if i%100 == 0 {
// 			log.Printf("Inserted %d rows so far...", i)
// 		}
// 	}

// 	duration := time.Since(start)
// 	log.Printf("Completed insertion of %d rows in %v\n", rowCount, duration)
// }

// // Generate a random 8-character short code
// func generateUniqueShortCode() string {
// 	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	code := make([]byte, 8)
// 	_, err := rand.Read(code)
// 	if err != nil {
// 		log.Fatalf("Failed to generate short code: %v", err)
// 	}
// 	for i := range code {
// 		code[i] = charset[code[i]%byte(len(charset))]
// 	}
// 	return string(code)
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"strings"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	// Connect to the database
// 	db, err := sql.Open("mysql", "root:Saibaba@2024@tcp(127.0.0.1:3306)/url_shortner")
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	defer db.Close()

// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("Database connection error: %v", err)
// 	}
// 	log.Println("Database connection successful!")

// 	// Replace these with actual short codes from your database
// 	shortCodes := []string{"0000ShIg", "0004YxzS", "0009gbt2"}

// 	// Fetch original URLs for selected short codes
// 	fetchOriginalURLs(db, shortCodes)
// }

// // Fetch original URLs for given short codes
// func fetchOriginalURLs(db *sql.DB, shortCodes []string) {
// 	// Generate query dynamically
// 	query := "SELECT original_url FROM url_shortner WHERE short_code IN (?" + strings.Repeat(", ?", len(shortCodes)-1) + ")"

// 	// Convert shortCodes to []interface{} for the query
// 	args := make([]interface{}, len(shortCodes))
// 	for i, code := range shortCodes {
// 		args[i] = code
// 	}

// 	// Measure the time taken to execute the query
// 	start := time.Now()
// 	rows, err := db.Query(query, args...)
// 	if err != nil {
// 		log.Fatalf("Failed to execute query: %v", err)
// 	}
// 	defer rows.Close()

// 	duration := time.Since(start)

// 	// Print fetched URLs
// 	fmt.Println("Fetched Original URLs:")
// 	for rows.Next() {
// 		var originalURL string
// 		if err := rows.Scan(&originalURL); err != nil {
// 			log.Fatalf("Failed to scan result: %v", err)
// 		}
// 		fmt.Println(originalURL)
// 	}

// 	// Log the time taken
// 	log.Printf("Time taken to fetch original URLs: %v\n", duration)
// }

package main

import (
	"database/sql"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:Saibaba@2024@tcp(127.0.0.1:3306)/url_shortner")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	log.Println("Database connection successful!")

	// Replace these with actual short codes from your database
	shortCodes := []string{"0000ShIg", "0004ChRr", "0005w7q9"}

	// Define the number of iterations (1 million times)
	iterations := 1000000

	// Start time before the loop
	start := time.Now()

	// Run the query 1 million times
	for i := 0; i < iterations; i++ {
		fetchOriginalURLs(db, shortCodes)

		// Optionally, log progress for every 100,000 iterations
		if i%100000 == 0 {
			log.Printf("Completed %d iterations...\n", i)
		}
	}

	// End time after the loop
	duration := time.Since(start)

	// Log the time taken for all iterations
	log.Printf("Completed fetching original URLs for %d iterations in %v\n", iterations, duration)
}

// Fetch original URLs for given short codes
func fetchOriginalURLs(db *sql.DB, shortCodes []string) {
	// Generate query dynamically
	query := "SELECT original_url FROM url_shortner WHERE short_code IN (?" + strings.Repeat(", ?", len(shortCodes)-1) + ")"

	// Convert shortCodes to []interface{} for the query
	args := make([]interface{}, len(shortCodes))
	for i, code := range shortCodes {
		args[i] = code
	}

	// Execute the query
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	defer rows.Close()

	// Print fetched URLs (optional)
	for rows.Next() {
		var originalURL string
		if err := rows.Scan(&originalURL); err != nil {
			log.Fatalf("Failed to scan result: %v", err)
		}
		// Print the original URL (optional)
		// fmt.Println(originalURL)
	}

	// No need to print out fetched URLs every iteration, just calculate duration
}

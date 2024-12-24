package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:Saibaba@2024@tcp(127.0.0.1:3306)/url_shortner")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Insert 1000 rows
	stmt, err := db.Prepare("INSERT INTO url_shortener (original_url, short_code) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 1; i <= 1000; i++ {
		originalURL := fmt.Sprintf("https://example.com/long-url-%d", i)
		shortCode := generateShortCode()
		_, err := stmt.Exec(originalURL, shortCode)
		if err != nil {
			log.Printf("Error inserting row %d: %v\n", i, err)
		}
	}

	fmt.Println("Inserted 1000 rows successfully!")
}

func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]rune, 8)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

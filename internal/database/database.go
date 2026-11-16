package database

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectWithRetry(dsn string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	maxRetries := 5            // Set the maximum number of retries
	backoff := 2 * time.Second // Set the initial backoff duration

	for attempt := 0; attempt <= maxRetries; attempt++ {
		log.Printf("Attempt %d of %d: Connecting to the database...", attempt+1, maxRetries)
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				db.SetMaxOpenConns(25) 
				db.SetMaxIdleConns(25)
				db.SetConnMaxLifetime(5 * time.Minute) 
				return db, nil // Connection successful
			}
		}
		log.Printf("Failed to connect to the database (attempt %d/%d): %v", attempt+1, maxRetries, err)
		time.Sleep(backoff)

		// Increase the backoff duration for the next retry
		backoff *= 2 // Exponential backoff

	}

	return nil, err // Return the last error after exhausting retries
}

package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yourusername/go-wallets/internal/logger"
)

func ConnectWithRetry(dsn string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	maxRetries := 5            // Set the maximum number of retries
	backoff := 2 * time.Second // Set the initial backoff duration

	for attempt := 0; attempt <= maxRetries; attempt++ {
		logger.Log.Info("Connecting to the database...", "attempt", attempt+1,   "of, maxRetries", maxRetries)
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
		logger.Log.Error("Failed to connect to the database", "attempt", attempt+1,   "of, maxRetries", maxRetries, "err", err)
		time.Sleep(backoff)

		// Increase the backoff duration for the next retry
		backoff *= 2 // Exponential backoff

	}

	return nil, err // Return the last error after exhausting retries
}

package middleware

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	// Open the connection
	postgresUrl := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", postgresUrl)

	if err != nil {
		log.Printf("Failed to connect to postgres on %s. Error: %v", postgresUrl, err)
		return nil, err
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		log.Printf("Failed to ping postgres on %s. Error: %v", postgresUrl, err)
		return nil, err
	}

	log.Printf("Successfully connected to postgres on %s", postgresUrl)
	// return the connection
	return db, nil
}

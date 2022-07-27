package middleware

import (
	"database/sql"
	"fmt"
	"github.com/Abashinos/otus_hw/server/util"
	_ "github.com/lib/pq"
	"log"
)

func CreateConnection() (*sql.DB, error) {
	// Open the connection
	postgresSafeUrl := fmt.Sprintf("postgresql://%s:%%s@%s:%s/%s?sslmode=%s",
		util.GetEnv("POSTGRES_USER", "postgres"),
		util.GetEnv("POSTGRES_SERVICE_HOST", "postgres"),
		util.GetEnv("POSTGRES_SERVICE_PORT", "5423"),
		util.GetEnv("POSTGRES_DB", "postgresdb"),
		util.GetEnv("POSTGRES_SSL_MODE", "disable"),
	)
	postgresUrl := fmt.Sprintf(
		postgresSafeUrl,
		util.GetEnv("POSTGRES_PASSWORD", "test123"),
	)
	db, err := sql.Open("postgres", postgresUrl)

	if err != nil {
		log.Printf("Failed to connect to postgres on %s. Error: %v", postgresSafeUrl, err)
		return nil, err
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		log.Printf("Failed to ping postgres on %s. Error: %v", postgresSafeUrl, err)
		return nil, err
	}

	log.Printf("Successfully connected to postgres on %s", postgresSafeUrl)
	// return the connection
	return db, nil
}

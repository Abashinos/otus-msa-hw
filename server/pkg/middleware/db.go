package middleware

import (
	"fmt"
	"github.com/Abashinos/otus-msa-hw/server/pkg"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func CreateConnection() (*gorm.DB, error) {
	// Open the connection
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		util.GetEnv("POSTGRES_SERVICE_HOST", "postgres"),
		util.GetEnv("POSTGRES_SERVICE_PORT", "5423"),
		util.GetEnv("POSTGRES_USER", "postgres"),
		util.GetEnv("POSTGRES_PASSWORD", "test123"),
		util.GetEnv("POSTGRES_DB", "postgresdb"),
		util.GetEnv("POSTGRES_SSL_MODE", "disable"),
	)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to postgres on %s. Error: %v", dsn, err)
		return nil, err
	}

	// check the connection
	db, err := conn.DB()
	if err != nil {
		log.Printf("Failed to get DB. Error: %v", err)
		return nil, err
	}
	defer db.Close()

	if db.Ping() != nil {
		log.Printf("Failed to ping postgres. Error: %v", err)
		return nil, err
	}

	log.Printf("Successfully connected to postgres")
	// return the connection
	return conn, nil
}

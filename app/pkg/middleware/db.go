package middleware

import (
	"fmt"
	"github.com/Abashinos/otus-msa-hw/app/pkg"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func CreateDBConnection() (*gorm.DB, error) {
	// Open the connection
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		util.GetEnv("POSTGRES_SERVICE_HOST", "postgres"),
		util.GetEnv("POSTGRES_SERVICE_PORT", "5432"),
		util.GetEnv("POSTGRES_USER", "postgres"),
		util.GetEnv("POSTGRES_PASSWORD", "test123"),
		util.GetEnv("POSTGRES_DB", "postgresdb"),
		util.GetEnv("POSTGRES_SSL_MODE", "disable"),
	)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect to postgres. Error: %v", err)
		return nil, err
	}

	log.Printf("Successfully connected to postgres")
	return conn, nil
}

func PingDB(dbConn *gorm.DB) error {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return err
	}

	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return err
	}
	return nil
}

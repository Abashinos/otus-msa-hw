package main

import (
	"fmt"
	"github.com/Abashinos/otus-msa-hw/server/pkg/middleware"
	"github.com/Abashinos/otus-msa-hw/server/pkg/models"
	"os"
)

func migrate() {
	db, err := middleware.CreateConnection()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}

func main() {
	//migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "up":
			migrate()
			os.Exit(0)
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			os.Exit(1)
		}
	}
}

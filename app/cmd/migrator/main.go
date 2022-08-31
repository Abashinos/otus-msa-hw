package main

import (
	"errors"
	"fmt"
	"github.com/Abashinos/otus-msa-hw/app/cmd/migrator/migrations"
	"github.com/Abashinos/otus-msa-hw/app/pkg/middleware"
	"github.com/Abashinos/otus-msa-hw/app/pkg/models"
	"github.com/akamensky/argparse"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func auto() (err error) {
	err = db.AutoMigrate(&models.User{})
	return
}

func up() (err error) {
	err = migrations.Up(db)
	return
}

var commands = map[string]func() error{
	"auto":   auto,
	"up":     up,
	"down":   func() error { return nil },
	"status": func() error { return nil },
}

func main() {
	parser := argparse.NewParser("Migrator", "It migrates")
	command := parser.StringPositional(&argparse.Options{Required: true, Validate: func(args []string) error {
		if _, ok := commands[args[0]]; !ok {
			return errors.New(fmt.Sprintf("Unsupported command '%s'", args[0]))
		}
		return nil
	}})

	var err error
	if err = parser.Parse(os.Args); err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	if db, err = middleware.CreateConnection(); err != nil {
		panic(err)
	}

	if err = commands[*command](); err != nil {
		panic(err)
	}
}

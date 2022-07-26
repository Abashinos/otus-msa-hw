package main

import (
	"github.com/Abashinos/otus_hw/server/middleware"
)

type AppHealth struct {
	Status string `json:"status"`
}

type DBHealth struct {
	OK    bool  `json:"ok"`
	Error error `json:"error"`
}

type Health struct {
	App *AppHealth `json:"app"`
	DB  *DBHealth  `json:"db"`
}

func dbHealth() *DBHealth {
	db, err := middleware.CreateConnection()
	if err != nil {
		return &DBHealth{
			OK:    false,
			Error: err,
		}
	}
	defer db.Close()
	return &DBHealth{
		OK:    true,
		Error: nil,
	}
}

func NewHealth() *Health {
	return &Health{
		App: &AppHealth{
			Status: "OK",
		},
		DB: dbHealth(),
	}
}

func checkHealth() *Health {
	return NewHealth()
}

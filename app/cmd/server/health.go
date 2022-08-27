package main

import (
	"github.com/Abashinos/otus-msa-hw/app/pkg/middleware"
	"gorm.io/gorm"
)

type AppHealth struct {
	Status string `json:"status"`
}

type DBHealth struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

type Health struct {
	App *AppHealth `json:"app"`
	DB  *DBHealth  `json:"db"`
}

func dbHealth(dbConn *gorm.DB) *DBHealth {
	if err := middleware.Ping(dbConn); err != nil {
		return &DBHealth{
			OK:    false,
			Error: err.Error(),
		}
	}
	return &DBHealth{
		OK:    true,
		Error: "",
	}
}

func checkHealth(dbConn *gorm.DB) *Health {
	return &Health{
		App: &AppHealth{
			Status: "OK",
		},
		DB: dbHealth(dbConn),
	}
}

package main

import (
	"api-golang-codebase/config"
	"api-golang-codebase/internal/app"
	"api-golang-codebase/internal/domain/event"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := app.NewDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close
	db.AutoMigrate(&event.Event{})

	server := app.NewServer(cfg, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/revpos/eTrak/internal/config"
	"github.com/revpos/eTrak/internal/handler"
	"github.com/revpos/eTrak/internal/repository"
	"github.com/revpos/eTrak/internal/service"
)

func main() {
	// Load the server configuration settings
	cfg := config.Load()

	// Setup/Open the database connection
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to DB: ", err)
	}
	defer db.Close()

	// Connect the architecture pipeline
	repo := repository.NewEventRepository(db)
	svc := service.NewEventService(repo)
	handler := handler.NewEventHandler(svc)

	// Root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the eTrak!")
	})

	// Track the new event
	http.HandleFunc("/track", handler.TrackEvent)

	log.Println("eTrak running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

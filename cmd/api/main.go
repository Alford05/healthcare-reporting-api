package main

import (
	"log"
	"net/http"

	"github.com/Alford05/healthcare-reporting-api/internal/db"

	"github.com/Alford05/healthcare-reporting-api/internal/handlers"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthHandler(conn))
	mux.HandleFunc("GET /api/patients", handlers.GetPatientsHandler(conn))
	mux.HandleFunc("GET /api/visits", handlers.GetVisitsHandler(conn))

	log.Println("server starting on port 8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/Alford05/healthcare-reporting-api/internal/db"
	"github.com/Alford05/healthcare-reporting-api/internal/handlers"
	"github.com/Alford05/healthcare-reporting-api/internal/logger"
	"github.com/Alford05/healthcare-reporting-api/internal/middleware"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	logr := logger.New()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthHandler(conn))
	mux.HandleFunc("GET /api/patients", handlers.GetPatientsHandler(conn))
	mux.HandleFunc("GET /api/visits", handlers.GetVisitsHandler(conn))
	mux.HandleFunc("GET /api/reports/therapist-productivity", handlers.TherapistProductivityReportHandler(conn))
	mux.HandleFunc("GET /api/reports/documentation-compliance", handlers.DocumentationComplianceReportHandler(conn))
	mux.HandleFunc("GET /api/reports/department-productivity", handlers.DepartmentProductivityReportHandler(conn))

	loggedMux := middleware.Logging(logr, mux)

	logr.Info("server starting", "port", 8080)

	err = http.ListenAndServe(":8080", loggedMux)
	if err != nil {
		log.Fatal(err)
	}
}

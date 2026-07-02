package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Alford05/healthcare-reporting-api/internal/repository"
)

func TherapistProductivityReportHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		report, err := repository.GetTherapistProductivityReport(db)
		if err != nil {
			log.Printf("therapist productivity report error: %v", err)

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "failed to fetch therapist productivity report",
			})
			return
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(report)
	}
}

func DocumentationComplianceReportHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		report, err := repository.GetDocumentationComplianceReport(db)
		if err != nil {
			log.Printf("documentation compliance report error: %v", err)

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "failed to fetch documentation compliance report",
			})
			return
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(report)
	}
}

func DepartmentProductivityReportHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		report, err := repository.GetDepartmentProductivityReport(db)
		if err != nil {
			log.Printf("department productivity report error: %v", err)

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "failed to fetch department productivity report",
			})
			return
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(report)
	}
}

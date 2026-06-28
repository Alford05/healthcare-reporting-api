package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Alford05/healthcare-reporting-api/internal/repository"
)

func GetPatientsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		patients, err := repository.GetPatients(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response := map[string]string{
				"error": "failed to fetch patients",
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(patients)
	}
}

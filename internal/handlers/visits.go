package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alford05/healthcare-reporting-api/internal/repository"
)

func GetVisitsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		filters := repository.VisitFilters{}

		therapistID := r.URL.Query().Get("therapist_id")
		if therapistID != "" {
			id, err := strconv.Atoi(therapistID)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "invalid therapist_id",
				})
				return
			}

			filters.TherapistID = id
		}
		filters.VisitType = r.URL.Query().Get("visit_type")
		filters.StartDate = r.URL.Query().Get("start_date")
		filters.EndDate = r.URL.Query().Get("end_date")

		visits, err := repository.GetVisits(db, filters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "failed to fetch visits",
			})
			return
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(visits)
	}
}

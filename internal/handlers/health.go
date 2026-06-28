package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func HealthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := db.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response := map[string]string{
				"status":   "error",
				"database": "disconnected",
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		response := map[string]string{
			"status":   "ok",
			"database": "connected",
		}

		json.NewEncoder(w).Encode(response)
	}
}

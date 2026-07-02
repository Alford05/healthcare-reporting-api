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

		page := 1
		pageSize := 10

		pageParam := r.URL.Query().Get("page")
		if pageParam != "" {
			parsedPage, err := strconv.Atoi(pageParam)
			if err != nil || parsedPage < 1 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "invalid page",
				})
				return
			}
			page = parsedPage
		}

		pageSizeParam := r.URL.Query().Get("page_size")
		if pageSizeParam != "" {
			parsedPageSize, err := strconv.Atoi(pageSizeParam)
			if err != nil || parsedPageSize < 1 {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "invalid page_size",
				})
				return
			}
			pageSize = parsedPageSize
		}

		filters.Page = page
		filters.PageSize = pageSize

		visits, err := repository.GetVisits(db, filters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "failed to fetch visits",
			})
			return
		}
		totalRecords, err := repository.CountVisits(db, filters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "failed to count visits",
			})
			return
		}

		totalPages := 0
		if pageSize > 0 {
			totalPages = (totalRecords + pageSize - 1) / pageSize
		}

		response := map[string]any{
			"data":          visits,
			"page":          page,
			"page_size":     pageSize,
			"total_records": totalRecords,
			"total_pages":   totalPages,
		}

		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		encoder.Encode(response)
	}
}

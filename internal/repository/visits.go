package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Alford05/healthcare-reporting-api/internal/models"
)

type VisitFilters struct {
	TherapistID int
	VisitType   string
	StartDate   string
	EndDate     string
	Page        int
	PageSize    int
}

func GetVisits(db *sql.DB, filters VisitFilters) ([]models.TherapyVisit, error) {
	query := `
		SELECT visit_id, patient_id, therapist_id, visit_date, visit_type, duration_minutes, notes_completed
		FROM therapy_visits
	`

	conditions := []string{}
	args := []any{}
	argPos := 1

	if filters.TherapistID != 0 {
		conditions = append(
			conditions,
			fmt.Sprintf("therapist_id = $%d", argPos),
		)

		args = append(args, filters.TherapistID)
		argPos++
	}

	if filters.VisitType != "" {
		conditions = append(
			conditions,
			fmt.Sprintf("visit_type = $%d", argPos),
		)

		args = append(args, filters.VisitType)
		argPos++
	}

	if filters.StartDate != "" {
		conditions = append(
			conditions,
			fmt.Sprintf("visit_date >= $%d", argPos),
		)

		args = append(args, filters.StartDate)
		argPos++
	}

	if filters.EndDate != "" {
		conditions = append(
			conditions,
			fmt.Sprintf("visit_date <= $%d", argPos),
		)

		args = append(args, filters.EndDate)
		argPos++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY visit_id"

	if filters.PageSize > 0 {
		offset := (filters.Page - 1) * filters.PageSize

		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
		args = append(args, filters.PageSize, offset)
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visits []models.TherapyVisit

	for rows.Next() {
		var visit models.TherapyVisit

		err := rows.Scan(
			&visit.VisitID,
			&visit.PatientID,
			&visit.TherapistID,
			&visit.VisitDate,
			&visit.VisitType,
			&visit.DurationMinutes,
			&visit.NotesCompleted,
		)
		if err != nil {
			return nil, err
		}

		visits = append(visits, visit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return visits, nil
}

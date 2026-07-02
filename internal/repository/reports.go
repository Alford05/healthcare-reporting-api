package repository

import (
	"database/sql"

	"github.com/Alford05/healthcare-reporting-api/internal/models"
)

func GetTherapistProductivityReport(db *sql.DB) ([]models.TherapistProductivityReport, error) {
	rows, err := db.Query(`
		SELECT
			t.therapist_id,
			t.first_name || ' ' || t.last_name AS therapist_name,
			t.discipline,
			COUNT(v.visit_id) AS total_visits,
			COALESCE(SUM(v.duration_minutes), 0) AS total_minutes
		FROM therapists t
		LEFT JOIN therapy_visits v
			ON t.therapist_id = v.therapist_id
		GROUP BY t.therapist_id, t.first_name, t.last_name, t.discipline
		ORDER BY total_visits DESC, t.therapist_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []models.TherapistProductivityReport

	for rows.Next() {
		var report models.TherapistProductivityReport

		err := rows.Scan(
			&report.TherapistID,
			&report.TherapistName,
			&report.Discipline,
			&report.TotalVisits,
			&report.TotalMinutes,
		)
		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func GetDocumentationComplianceReport(db *sql.DB) ([]models.DocumentationComplianceReport, error) {
	rows, err := db.Query(`
		SELECT
			t.therapist_id,
			t.first_name || ' ' || t.last_name AS therapist_name,
			t.discipline,
			COUNT(dm.metric_id) AS total_notes,
			COALESCE(SUM(CASE WHEN dm.is_late = true THEN 1 ELSE 0 END), 0) AS late_notes,
			ROUND(
				(
					COUNT(dm.metric_id) - COALESCE(SUM(CASE WHEN dm.is_late = true THEN 1 ELSE 0 END), 0)
				)::numeric / NULLIF(COUNT(dm.metric_id), 0) * 100,
				2
			) AS compliance_percentage
		FROM therapists t
		LEFT JOIN therapy_visits v
			ON t.therapist_id = v.therapist_id
		LEFT JOIN documentation_metrics dm
			ON v.visit_id = dm.visit_id
		GROUP BY t.therapist_id, t.first_name, t.last_name, t.discipline
		ORDER BY compliance_percentage DESC NULLS LAST, t.therapist_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []models.DocumentationComplianceReport

	for rows.Next() {
		var report models.DocumentationComplianceReport

		err := rows.Scan(
			&report.TherapistID,
			&report.TherapistName,
			&report.Discipline,
			&report.TotalNotes,
			&report.LateNotes,
			&report.CompliancePercentage,
		)
		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func GetDepartmentProductivityReport(db *sql.DB) ([]models.DepartmentProductivityReport, error) {
	rows, err := db.Query(`
		SELECT
			d.department_id,
			d.name AS department_name,
			COUNT(v.visit_id) AS total_visits,
			COALESCE(SUM(v.duration_minutes), 0) AS total_minutes,
			COALESCE(
				ROUND(
					AVG(v.duration_minutes)::numeric,
					2
				),
				0
			) AS avg_minutes_per_visit
		FROM departments d
		LEFT JOIN therapists t
			ON d.department_id = t.department_id
		LEFT JOIN therapy_visits v
			ON t.therapist_id = v.therapist_id
		GROUP BY d.department_id, d.name
		ORDER BY total_visits DESC, d.department_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []models.DepartmentProductivityReport

	for rows.Next() {
		var report models.DepartmentProductivityReport

		err := rows.Scan(
			&report.DepartmentID,
			&report.DepartmentName,
			&report.TotalVisits,
			&report.TotalMinutes,
			&report.AvgMinutesPerVisit,
		)
		if err != nil {
			return nil, err
		}

		reports = append(reports, report)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

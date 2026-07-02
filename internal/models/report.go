package models

type TherapistProductivityReport struct {
	TherapistID   int    `json:"therapist_id"`
	TherapistName string `json:"therapist_name"`
	Discipline    string `json:"discipline"`
	TotalVisits   int    `json:"total_visits"`
	TotalMinutes  int    `json:"total_minutes"`
}

type DocumentationComplianceReport struct {
	TherapistID          int     `json:"therapist_id"`
	TherapistName        string  `json:"therapist_name"`
	Discipline           string  `json:"discipline"`
	TotalNotes           int     `json:"total_notes"`
	LateNotes            int     `json:"late_notes"`
	CompliancePercentage float64 `json:"compliance_percentage"`
}

type DepartmentProductivityReport struct {
	DepartmentID       int     `json:"department_id"`
	DepartmentName     string  `json:"department_name"`
	TotalVisits        int     `json:"total_visits"`
	TotalMinutes       int     `json:"total_minutes"`
	AvgMinutesPerVisit float64 `json:"avg_minutes_per_visit"`
}

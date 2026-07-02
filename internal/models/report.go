package models

type TherapistProductivityReport struct {
	TherapistID   int    `json:"therapist_id"`
	TherapistName string `json:"therapist_name"`
	Discipline    string `json:"discipline"`
	TotalVisits   int    `json:"total_visits"`
	TotalMinutes  int    `json:"total_minutes"`
}

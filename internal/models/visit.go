package models

type TherapyVisit struct {
	VisitID         int    `json:"visit_id"`
	PatientID       int    `json:"patient_id"`
	TherapistID     int    `json:"therapist_id"`
	VisitDate       string `json:"visit_date"`
	VisitType       string `json:"visit_type"`
	DurationMinutes int    `json:"duration_minutes"`
	NotesCompleted  bool   `json:"notes_completed"`
}

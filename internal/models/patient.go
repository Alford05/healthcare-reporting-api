package models

type Patient struct {
	PatientID     int     `json:"patient_id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Age           int     `json:"age"`
	Sex           string  `json:"sex"`
	Diagnosis     string  `json:"diagnosis"`
	AdmitDate     string  `json:"admit_date"`
	DischargeDate *string `json:"discharge_date"`
}

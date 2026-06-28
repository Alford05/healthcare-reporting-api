package repository

import (
	"database/sql"

	"github.com/Alford05/healthcare-reporting-api/internal/models"
)

func GetPatients(db *sql.DB) ([]models.Patient, error) {
	rows, err := db.Query(`
		SELECT patient_id, first_name, last_name, age, sex, diagnosis, admit_date, discharge_date
		FROM patients
		ORDER BY patient_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient

	for rows.Next() {
		var patient models.Patient

		err := rows.Scan(
			&patient.PatientID,
			&patient.FirstName,
			&patient.LastName,
			&patient.Age,
			&patient.Sex,
			&patient.Diagnosis,
			&patient.AdmitDate,
			&patient.DischargeDate,
		)
		if err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}

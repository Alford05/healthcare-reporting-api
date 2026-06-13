CREATE INDEX idx_patients_diagnosis
ON patients(diagnosis);

CREATE INDEX idx_patients_admit_date
ON patients(admit_date);

CREATE INDEX idx_therapists_department_id
ON therapists(department_id);

CREATE INDEX idx_therapy_visits_patient_id
ON therapy_visits(patient_id);

CREATE INDEX idx_therapy_visits_therapist_id
ON therapy_visits(therapist_id);

CREATE INDEX idx_therapy_visits_visit_date
ON therapy_visits(visit_date);

CREATE INDEX idx_therapy_visits_visit_type
ON therapy_visits(visit_type);

CREATE INDEX idx_therapy_visits_type_date
ON therapy_visits(visit_type, visit_date);

CREATE INDEX idx_documentation_metrics_visit_id
ON documentation_metrics(visit_id);

CREATE INDEX idx_documentation_metrics_is_late
ON documentation_metrics(is_late);

CREATE INDEX idx_staffing_department_date
ON staffing(department_id, staffing_date);

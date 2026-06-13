CREATE TABLE departments (
    department_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE patients (
    patient_id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INT NOT NULL CHECK (age >= 0),
    sex TEXT NOT NULL,
    diagnosis TEXT NOT NULL,
    admit_date DATE NOT NULL,
    discharge_date DATE
);

CREATE TABLE therapists (
    therapist_id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    discipline TEXT NOT NULL,
    department_id INT NOT NULL REFERENCES departments(department_id)
);

CREATE TABLE therapy_visits (
    visit_id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL REFERENCES patients(patient_id),
    therapist_id INT NOT NULL REFERENCES therapists(therapist_id),
    visit_date DATE NOT NULL,
    visit_type TEXT NOT NULL,
    duration_minutes INT NOT NULL CHECK (duration_minutes > 0),
    notes_completed BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE documentation_metrics (
    metric_id SERIAL PRIMARY KEY,
    visit_id INT NOT NULL UNIQUE REFERENCES therapy_visits(visit_id) ON DELETE CASCADE,
    note_signed_at TIMESTAMP,
    note_due_at TIMESTAMP NOT NULL,
    is_late BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE staffing (
    staffing_id SERIAL PRIMARY KEY,
    department_id INT NOT NULL REFERENCES departments(department_id),
    staffing_date DATE NOT NULL,
    scheduled_therapists INT NOT NULL CHECK (scheduled_therapists >= 0),
    actual_therapists INT NOT NULL CHECK (actual_therapists >= 0)
);

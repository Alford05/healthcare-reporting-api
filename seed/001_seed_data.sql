INSERT INTO departments (name)
VALUES
('Physical Therapy'),
('Occupational Therapy'),
('Speech Therapy');

INSERT INTO patients (
    first_name,
    last_name,
    age,
    sex,
    diagnosis,
    admit_date,
    discharge_date
)
VALUES
('John', 'Smith', 72, 'M', 'Stroke', '2026-01-02', '2026-01-20'),
('Jane', 'Doe', 81, 'F', 'Hip Fracture', '2026-01-05', '2026-01-25'),
('Robert', 'Johnson', 67, 'M', 'Pneumonia', '2026-01-10', NULL),
('Emily', 'Williams', 75, 'F', 'CHF', '2026-01-12', NULL),
('Michael', 'Brown', 43, 'M', 'TBI', '2026-01-15', NULL),
('Sarah', 'Davis', 69, 'F', 'CVA', '2026-01-18', NULL);

INSERT INTO therapists (
    first_name,
    last_name,
    discipline,
    department_id
)
VALUES
('Lisa', 'Miller', 'PT', 1),
('David', 'Wilson', 'PT', 1),
('Karen', 'Moore', 'OT', 2),
('James', 'Taylor', 'OT', 2),
('Jennifer', 'Anderson', 'SLP', 3);

INSERT INTO therapy_visits (
    patient_id,
    therapist_id,
    visit_date,
    visit_type,
    duration_minutes,
    notes_completed
)
VALUES
(1,1,'2026-01-03','Evaluation',60,true),
(1,1,'2026-01-05','Treatment',45,true),
(1,1,'2026-01-07','Treatment',45,true),

(2,2,'2026-01-06','Evaluation',60,true),
(2,2,'2026-01-08','Treatment',40,true),
(2,2,'2026-01-10','Treatment',40,false),

(3,3,'2026-01-11','Evaluation',55,true),
(3,3,'2026-01-13','Treatment',35,true),

(4,4,'2026-01-13','Evaluation',60,true),
(4,4,'2026-01-15','Treatment',40,false),

(5,1,'2026-01-16','Evaluation',60,true),
(5,1,'2026-01-18','Treatment',45,true),

(6,5,'2026-01-19','Evaluation',50,true),
(6,5,'2026-01-21','Treatment',30,true);

INSERT INTO documentation_metrics (
    visit_id,
    note_signed_at,
    note_due_at,
    is_late
)
VALUES
(1,'2026-01-03 10:00','2026-01-03 23:59',false),
(2,'2026-01-05 11:00','2026-01-05 23:59',false),
(3,'2026-01-08 08:00','2026-01-07 23:59',true),
(4,'2026-01-06 12:00','2026-01-06 23:59',false),
(5,'2026-01-08 14:00','2026-01-08 23:59',false),
(6,NULL,'2026-01-10 23:59',true),
(7,'2026-01-11 10:00','2026-01-11 23:59',false),
(8,'2026-01-13 10:00','2026-01-13 23:59',false),
(9,'2026-01-13 12:00','2026-01-13 23:59',false),
(10,NULL,'2026-01-15 23:59',true),
(11,'2026-01-16 09:00','2026-01-16 23:59',false),
(12,'2026-01-18 10:00','2026-01-18 23:59',false),
(13,'2026-01-19 11:00','2026-01-19 23:59',false),
(14,'2026-01-21 13:00','2026-01-21 23:59',false);

INSERT INTO staffing (
    department_id,
    staffing_date,
    scheduled_therapists,
    actual_therapists
)
VALUES
(1,'2026-01-01',5,5),
(1,'2026-01-02',5,4),
(1,'2026-01-03',5,5),

(2,'2026-01-01',4,4),
(2,'2026-01-02',4,3),
(2,'2026-01-03',4,4),

(3,'2026-01-01',2,2),
(3,'2026-01-02',2,2),
(3,'2026-01-03',2,1);


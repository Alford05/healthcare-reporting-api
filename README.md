# Healthcare Reporting API

A Go-based backend reporting service built with Go, PostgreSQL, and Docker that simulates analytics workflows commonly used in healthcare organizations. 

The application exposes REST API endpoints for patient management, therapy visit tracking, operational reporting, and performance analytics. The project demonstrates backend engineering concepts including SQL query design, filtering, pagination, structured logging, Dockerized deployment, and layered application architecture.

---

## 🚀 Tech Stack

- Go (net/http)
- PostgreSQL
- Docker / Docker Compose
- REST APIs
- SQL (joins, aggregations, filtering)
- Structured JSON Logging
- Environment-based configuration 

---

## 📦 Features

### Core API
- GET `/health` – service + database health check
- GET `/api/patients` – patient listing
- GET `/api/visits` – therapy visit records

### Filtering
- Filter visits by:
  - therapist_id
  - visit_type
  - date range

### Pagination
- Page-based pagination on visit endpoints:
  - page
  - page_size
- Returns:
  - total_records
  - total_pages

### Reporting Endpoints
- Therapist productivity report
- Documentation compliance report
- Department productivity report

These reports are powered by SQL joins, aggregations, grouping, and calculated metrics to simulate real-world healthcare analytics workloads.

🧱 Architecture
                ┌─────────────────┐
                │     Client      │
                │ (curl / browser)│
                └────────┬────────┘
                         │
                         ▼
                ┌─────────────────┐
                │   HTTP Router   │
                │  + Middleware   │
                └────────┬────────┘
                         │
                         ▼
                ┌─────────────────┐
                │    Handlers     │
                │ Request/Response│
                └────────┬────────┘
                         │
                         ▼
                ┌─────────────────┐
                │   Repository    │
                │ SQL Data Access │
                └────────┬────────┘
                         │
                         ▼
                ┌─────────────────┐
                │   PostgreSQL    │
                │    Database     │
                └─────────────────┘
Project Structure
cmd/api             Application entry point
internal/db         Database connection logic
internal/handlers   HTTP handlers
internal/repository SQL queries and data access
internal/models     Domain models
migrations          Database schema
seed                Sample data
---

## 🐘 Database Design

The system uses PostgreSQL tables representing:

- patients
- therapists
- therapy_visits
- documentation_metrics
- departments

Relationships simulate a hospital therapy workflow:
Departments → therapists → Therapy Visits → patients → Therapy Visits → Documentation Metrics 

---

## 🐳 Running the Project

### 1. Start services
```bash
docker compose up --build
```

### 2. API runs on:
```bash
http://localhost:8080
```

### 3. Example requests
```bash
curl http://localhost:8080/health

curl "http://localhost:8080/api/visits?page=1&page_size=5"

curl "http://localhost:8080/api/reports/therapist-productivity"
```
### 📊 Example Response
{
  "data": [
    {
      "visit_id": 1,
      "patient_id": 1,
      "therapist_id": 1,
      "visit_type": "Treatment",
      "duration_minutes": 45
    }
  ],
  "page": 1,
  "page_size": 5,
  "total_records": 14,
  "total_pages": 3
}

🎯 Key Skills Demonstrated
Building REST APIs in Go using net/http
Designing layered backend architectures
Creating repository-based data access patterns
Writing complex SQL joins and aggregation queries
Implementing filtering and pagination
Developing reporting-focused database queries
Containerizing services with Docker and Docker Compose
Managing environment-based application configuration
Implementing structured request logging

📌 Future Improvements
JWT Authentication & Authorization
Redis Caching Layer
Unit & Integration Testing
OpenAPI / Swagger Documentation
CI/CD Pipeline with GitHub Actions
Database Migration Automation

👤 Author

Daniel Alford

Backend-focused software developer with a clinical healthcare background, building backend systems with Go, PostgreSQL, Docker, and cloud-native development practices.




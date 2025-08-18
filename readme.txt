📚 Go Gin Bookstore API

A Bookstore REST API built with Go (Gin framework) and PostgreSQL.
This project demonstrates secure authentication, clean project structure, API testing, and modern observability practices using Grafana, Prometheus, Loki, and InfluxDB.

✨ Features

User Management

Register new users with bcrypt password hashing

Secure login with JWT access & refresh tokens

Role-based user control (admin/user)

Bookstore Operations

Create, read, update, and delete books

Nested JSON fields: Publisher, Tags, Ratings

Authentication & Authorization

JWT-based access token & refresh flow

Middleware-protected routes

Monitoring & Observability

Prometheus & Grafana dashboards

Loki + Promtail for centralized log management

InfluxDB & VictoriaMetrics for metrics storage

Stress testing with k6

Security & Configuration

Environment variables managed via .env

Docker secrets for sensitive credentials

🛠 Tech Stack

Backend: Go + Gin

Database: PostgreSQL

Authentication: JWT + bcrypt

Monitoring: Prometheus, Grafana, Loki, InfluxDB, VictoriaMetrics

Logging: Loki + Promtail

Load Testing: k6

Containerization: Docker & Docker Compose

🚀 Getting Started
1️⃣ Clone the repository
git clone https://github.com/your-username/bookstore-api.git
cd bookstore-api

2️⃣ Configure Environment

Create a .env file:

POSTGRES_USER=bookstore
POSTGRES_PASSWORD=secret
POSTGRES_DB=bookstore_db
POSTGRES_HOST=postgres
POSTGRES_PORT=5432

JWT_SECRET=supersecretjwt

3️⃣ Run with Docker Compose
docker-compose up --build


Services started:

API → http://localhost:8080

Grafana → http://localhost:3000 (default user/pass: admin/admin)

Prometheus → http://localhost:9090

Loki → http://localhost:3100

InfluxDB → http://localhost:8086

📖 API Endpoints
🔑 Authentication

POST /register → Register new user

POST /login → Authenticate & receive JWT

POST /refresh → Get new access token from refresh token

📚 Books

POST /book → Create book (admin only)

GET /books → List all books

GET /books/:id → Get book by ID

PUT /books/:id → Update book (admin only)

DELETE /books/:id → Delete book (admin only)

📊 Monitoring
Grafana Dashboards

API Performance (via Prometheus & InfluxDB)

System Metrics (via Node Exporter)

Logs (via Loki + Promtail)

Example Setup
# Import k6 dashboard into Grafana
curl -X POST http://localhost:3000/api/dashboards/import \
  -H "Content-Type: application/json" \
  -u admin:admin \
  -d @grafana-dashboards/k6.json

🧪 Load / Stress Testing
Run k6 Tests
k6 run --out influxdb=http://localhost:8086/k6 scripts/stress_test.js

Example k6 Stages
export let options = {
  stages: [
    { duration: '30s', target: 50 },   // ramp-up
    { duration: '1m', target: 100 },   // steady load
    { duration: '1m', target: 100 },   // hold load
    { duration: '30s', target: 0 },    // ramp-down
  ],
};


Results can be visualized in Grafana → k6 Dashboard.

📜 Logs

Application logs are collected using Promtail and sent to Loki, which can be queried in Grafana using LogQL.

Example LogQL query:

{job="gin-app"} |= "ERROR"

🧩 Project Structure
├── internal/
│   ├── api/handlers/     # Route handlers
│   ├── domain/models/    # Database models
│   └── utils/            # Helper functions (JWT, hashing)
├── database/             # DB connection & migrations
├── docker-compose.yml    # Infra services
├── prometheus/           # Prometheus config
├── promtail/             # Promtail config
└── scripts/              # k6 load test scripts

✅ To Do / Roadmap

 Add unit & integration tests

 CI/CD pipeline with GitHub Actions

 API documentation with Swagger

 Horizontal scaling with Kubernetes

📜 License

MIT License. Free to use and modify.

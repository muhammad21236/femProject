# 🏋️‍♂️ Workout Tracker API (Go + PostgreSQL)

A practical workout tracker REST API built with Go, PostgreSQL, and the `chi` router package. This project was created as part of a professional Go course to learn how to build a complete multi-tiered HTTP server from scratch, with database integration, secure authentication, and robust unit testing.

---

## 📖 Features

- 📦 Build a complete HTTP server from scratch using the `chi` package
- 📊 Connect to a PostgreSQL database running in Docker
- 🔄 Implement database migrations using `goose`
- 📝 Design and build comprehensive CRUD API endpoints
- 🔒 Secure user authentication with password hashing and JWT
- 🛡️ Middleware for protecting routes and validating user ownership
- 🧪 Write and run unit tests with a dedicated test database
- 📚 Apply clean code structure and Go best practices

---

## 🗂️ Project Structure

```
.
├── internal/
│   ├── api/
│   │   ├── token_handler.go
│   │   ├── user_handler.go
│   │   ├── workout_handler.go
│   ├── app/
│   │   └── app.go
│   ├── middleware/
│   │   └── middleware.go
│   ├── routes/
│   │   └── routes.go
│   ├── store/
│   │   └── database.go
│   │   └── tokens.go
│   │   └── user_store.go
│   │   └── workout_store_test.go
│   │   └── workout_store.go
│   ├── tokens/
│   │   └── tokens.go
│   ├── utils/
│   │   └── utils.go
├── migrations/
│   └── 00001_users.sql
│   └── 00002_workouts.sql
│   └── 00003_workout_entries.sql
│   └── 00004_tokens.sql
│   └── fs.go
├──.gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## 🛠️ Getting Started

### 📦 Prerequisites

- Go 1.22+
- Docker & Docker Compose
- PostgreSQL (via Docker)
- `goose` for database migrations

---

### 🐳 Run with Docker Compose

```bash
docker-compose up --build
```

---

### 🐦 Run Migrations

```bash
goose -dir ./migrations postgres "postgresql://postgres:password@localhost:5433/workoutdb?sslmode=disable" up
```

---

### 🚀 Start the Server

```bash
go run main.go
```

---

## 📚 API Endpoints

### Workouts

| Method     | Endpoint                  | Description             |
| :--------- | :------------------------ | :---------------------- |
| `GET`    | `/workouts/{workoutID}` | Get workout by ID       |
| `POST`   | `/workouts`             | Create new workout      |
| `PUT`    | `/workouts/{workoutID}` | Update existing workout |
| `DELETE` | `/workouts/{workoutID}` | Delete workout          |

### Users

| Method   | Endpoint                   | Description                   |
| :------- | :------------------------- | :---------------------------- |
| `POST` | `/users`                 | Register new user             |
| `POST` | `/tokens/authentication` | Authenticate user, return JWT |

---

## 🔐 Authentication

- Passwords hashed with bcrypt
- JWT issued on successful login
- Middleware to validate and protect authenticated routes

---

## 🧪 Tests

- Unit tests using Go's testing package
- Isolated test database using Docker
- Test coverage for core business logic and handlers

---

## 📦 Technologies

- Go (golang.org)
- PostgreSQL
- Docker & Docker Compose
- `chi` router
- `pgx` PostgreSQL driver
- `goose` for migrations
- JWT for authentication

---

## 📖 Course Topics Covered

✅ Master Go fundamentals
✅ HTTP server setup and routing
✅ PostgreSQL integration via Docker
✅ Database migrations and schema design
✅ CRUD API implementation
✅ Secure user authentication
✅ Middleware creation
✅ Unit testing with isolated test database
✅ Clean, scalable Go application architecture

---

## 📃 License

This project is for educational and portfolio use.

---

## 🙌 Author

**Muhammad** | [GitHub](https://github.com/muhammad21236)

---

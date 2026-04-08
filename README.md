# Auth API

A REST API built with Go and Gin framework, featuring JWT authentication.

## Tech Stack
- Go
- Gin
- JWT
- PostgreSQL
- GORM

## Project Structure
```
auth-api/
├── main.go
├── handler/
│   └── auth.go      # Handle HTTP requests (register, login)
├── service/
│   └── auth.go      # Business logic
├── repository/
│   └── user.go      # Database operations
└── model/
    └── user.go      # Data structures
```

## Getting Started
```bash
go mod download
go run main.go
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/auth/register | Register a new user |
| POST | /api/auth/login | Login and get JWT token |

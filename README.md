# TinyRun

A simple running tracker application with Go backend and Flutter frontend.

## Backend

### Prerequisites
- Go 1.19 or higher
- SQLite

### Project Structure
```
tinyrun/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes.go
│   ├── config/
│   ├── database/
│   └── service/
├── models/
├── pkg/
│   └── utils/
├── go.mod
└── go.sum
```

### Setup
1. Clone the repository

2. Install dependencies
```bash
go mod tidy
```

3. Run the server
```bash
go run cmd/server/main.go
```

The server will start on `localhost:8080`

### API Endpoints

#### Authentication
- `POST /register` - Register a new user
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```

- `POST /login` - Login and receive API key
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```

- `GET /verify` - Verify API key (requires Authorization header)

#### Runs
All run endpoints require Authorization header with API key

- `POST /runs` - Save a new run
  ```json
  {
    "start_time": "2023-09-20T10:00:00Z",
    "end_time": "2023-09-20T10:30:00Z",
    "distance": 5000.0,
    "duration": 1800,
    "gps_points": "[{\"lat\":1.23,\"lng\":4.56},...]"
  }
  ```

- `GET /runs` - Get all runs for authenticated user

### Authentication
All protected endpoints require an API key in the Authorization header:
```
Authorization: YOUR_API_KEY
```

## Frontend (Flutter)

Android app for tinyrun is currently in development, concurrently with this project

#!/bin/bash

# Create directories
mkdir -p cmd/server
mkdir -p internal/api/handlers
mkdir -p internal/api/middleware
mkdir -p internal/config
mkdir -p internal/database
mkdir -p internal/service
mkdir -p models
mkdir -p pkg/utils

# Create files
touch cmd/server/main.go
touch internal/api/handlers/auth.go
touch internal/api/handlers/runs.go
touch internal/api/middleware/auth.go
touch internal/api/routes.go
touch internal/config/config.go
touch internal/database/db.go
touch internal/service/auth.go
touch internal/service/runs.go
touch models/user.go
touch models/run.go
touch pkg/utils/crypto.go
touch go.mod
touch go.sum

echo "Directory structure created successfully!"
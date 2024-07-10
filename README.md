# API-Template
API is a Go-based backend service for managing user registration and authentication. It uses PostgreSQL for database management and Redis for caching.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)

## Prerequisites

- Go 1.16+
- PostgreSQL
- Redis
- Git

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/isteportal-api.git
    cd isteportal-api
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Configuration

1. Update the database configuration in `config/database.go`:

    ```go
    db, err := sql.Open("postgres", "user=yourusername password=yourpassword dbname=yourdbname sslmode=disable")
    ```

2. Update the Redis configuration in `config/database.go`:

    ```go
    redisClient := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    ```

## Running the Application

To start the server, run:

```sh
go run main.go

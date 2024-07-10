# API-Template

This template is designed as a Go-based backend service. The application uses PostgreSQL for database management and Redis for caching. The project aims to implement a layered architecture to improve modularity and scalability and make it suitable for large-scale applications. Endponids are included as examples.
## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- Go 1.16+
- PostgreSQL
- Redis
- Git

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/dzdrgl/API-Template.git
    cd API-Template
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Configuration

1. Update the database configuration in `config/database.go`:

    ```go
    db, err := sql.Open("postgres", "user=yourUsername password=yourPassword dbname=your_DB sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to PostgreSQL: %v", err)
    }
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
    ```

The server will start on port `8080`. You can change the port by modifying the `Addr` field in the `main.go` file.

## API Endpoints

### User Endpoints

- `POST /api/v1/user/register` - Register a new user.
- `POST /api/v1/user/login` - Login an existing user.

### Middleware

- `AuthMiddleware` - Middleware to handle authentication.

## Project Structure

    ```plaintext
    api-template/
    ├── config/
    │   ├── database.go
    │   ├── handlers.go
    │   ├── repositories.go
    │   ├── routes.go
    │   └── services.go
    ├── handlers/
    │   └── user_handler.go
    ├── repositories/
    │   └── user_repository.go
    ├── services/
    │   └── user_service.go
    ├── main.go
    └── go.mod
    ```

- **config/**: Contains configuration and initialization of database, repositories, services, handlers, and routes.
- **handlers/**: Contains HTTP handlers.
- **repositories/**: Contains data access logic.
- **services/**: Contains business logic.
- **main.go**: The entry point of the application.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

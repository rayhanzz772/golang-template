# Golang Template

A starter template for building REST APIs in Go using [Fiber](https://gofiber.io/), [GORM](https://gorm.io/) (PostgreSQL), and JWT authentication.

## Features

- Layered architecture: `router` → `controller` → `service` → `repository`
- Modular structure (`modules/auth`, `modules/user`)
- JWT authentication middleware
- Request validation with `go-playground/validator`
- Standardized success/error response helpers with pagination
- Environment-based configuration via `.env`
- Graceful shutdown
- SQL migration files

## Tech Stack

| Purpose      | Library                                  |
| ------------ | ---------------------------------------- |
| HTTP server  | `github.com/gofiber/fiber/v2`            |
| ORM          | `gorm.io/gorm` + `gorm.io/driver/postgres` |
| Auth         | `github.com/golang-jwt/jwt/v5`           |
| Validation   | `github.com/go-playground/validator/v10` |
| Config       | `github.com/joho/godotenv`               |
| Password     | `golang.org/x/crypto/bcrypt`             |

## Requirements

- Go **1.25+** — https://go.dev/dl/
- PostgreSQL **13+** — https://www.postgresql.org/download/
- (Optional) `golang-migrate` CLI for running migrations — https://github.com/golang-migrate/migrate

## Installation

1. **Clone the repository**

   ```bash
   git clone <your-repo-url> golang-template
   cd golang-template
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Create the database** (if not already created)

   ```bash
   createdb -U postgres your_db
   # or: psql -U postgres -c "CREATE DATABASE your_db;"
   ```

4. **Configure environment variables**

   ```bash
   cp .env.example .env
   ```

   Then edit `.env` with your values:

   | Variable      | Description                                   | Default          |
   | ------------- | --------------------------------------------- | ---------------- |
   | `DB_HOST`     | PostgreSQL host                               | `localhost`      |
   | `DB_PORT`     | PostgreSQL port                               | `5432`           |
   | `DB_USER`     | PostgreSQL user                               | `postgres`       |
   | `DB_PASSWORD` | PostgreSQL password                           | `postgres`       |
   | `DB_NAME`     | Database name                                 | `your_db`        |
   | `SERVER_HOST` | Bind address for the HTTP server              | `0.0.0.0`        |
   | `PORT`        | HTTP port                                     | `8000`           |
   | `JWT_SECRET`  | Secret key used to sign JWTs (**required**)   | —                |
   | `JWT_EXPIRY`  | Token lifetime in hours                       | `24`             |

   > Generate a strong secret for production:
   > `openssl rand -hex 32`

## Running the Application

Development (with auto-reload via [air](https://github.com/air-verse/air)):

```bash
go install github.com/air-verse/air@latest
air
```

Or run directly:

```bash
go run .
```

Build a binary and run it:

```bash
go build -o bin/server .
./bin/server
```

The server starts on `http://localhost:8000` by default.

> On startup the app calls `AutoMigrate` for the `User` model. For production, prefer running the SQL migrations in `migrations/` instead.

## Database Migrations

Migration files live in `migrations/` and follow the `golang-migrate` naming convention (`<timestamp>_<name>.up.sql` / `.down.sql`).

Using the `migrate` CLI:

```bash
# Apply all migrations
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/your_db?sslmode=disable" up

# Roll back the last migration
migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/your_db?sslmode=disable" down 1
```

## API Overview

### Public

| Method | Endpoint       | Description          |
| ------ | -------------- | -------------------- |
| GET    | `/status`      | Health check (200 OK) |
| POST   | `/auth/login`  | Login and get a JWT  |

> There is no public register endpoint yet. Create the first user via `POST /api/users` after obtaining a token, or insert one using the seeder.

### Protected (requires `Authorization: Bearer <token>`)

| Method | Endpoint         | Description      |
| ------ | ---------------- | ---------------- |
| GET    | `/api/users`     | List users (paginated) |
| POST   | `/api/users`     | Create user      |
| GET    | `/api/users/:id` | Get user by ID   |
| PUT    | `/api/users/:id` | Update user      |
| DELETE | `/api/users/:id` | Delete user      |

Example login:

```bash
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'
```

## Project Structure

```
.
├── config/        # Environment loading and database connection
├── middleware/    # JWT authentication middleware
├── migrations/    # SQL migration files (up/down)
├── model/         # GORM models
├── modules/       # Feature modules (auth, user)
│   ├── auth/
│   └── user/
├── router/        # Route registration
├── seeder/        # Database seeders
├── utils/         # JWT, response, pagination, validation helpers
├── main.go        # Application entry point
└── .env.example   # Example environment variables
```

## Seeding the Database

`seeder/seed.go` contains `SeedUsers()`. Uncomment the call in `main.go`:

```go
seed.SeedUsers()
```

Then run the app once. It creates a test user (`john@example.com` / `password123`).

## Testing & Linting

```bash
go test ./...
go vet ./...
```

(No automated tests are included yet.)

# AGENTS.md

Go (1.25) REST API template: Fiber + GORM (PostgreSQL) + JWT. No CI, no tests, no lint config — verify with `go vet ./...` and `go build ./...`. Commit messages follow Conventional Commits (`feat: ...`).

## Commands

Single CLI with subcommands (default `server`):

```bash
go run . server      # start HTTP server on :8000
go run . seed        # seed sample user (idempotent)
go run . migrate     # GORM AutoMigrate of User model
make run|build|seed|migrate   # shortcuts; build outputs bin/app
air -- server        # hot reload (air takes subcommand as arg: `air -- server`)
```

Verify: `go build ./...` and `go vet ./...`.

## Gotchas

- Module path is `belajar-coding/go` (not derivable from the repo name).
- `main()` connects to PostgreSQL before dispatching **every** command — `seed`, `migrate`, and `server` all fail without a reachable DB and `JWT_SECRET` (hard fatal if missing). `.env` is gitignored: copy `.env.example` first and set `DB_*` (unset vars become empty, no validation).
- The `migrate` command is just `AutoMigrate(&model.User{})`, and server startup runs it on every boot. The SQL files in `migrations/` use `golang-migrate` naming but are **dead/unused** — nothing in the app reads them, and `go run . migrate` does not apply them. Don't confuse the two; the SQL also sets up `pgcrypto` and differs from AutoMigrate.
- There is no public register endpoint. First user comes from `go run . seed` (`john@example.com` / `password123`) or `POST /api/users` with a token.

## Structure

- Entry point `main.go` → `config/` (env + DB) → `router/router.go` (`SetupRoutes`).
- Feature modules live in `modules/<name>/` with `controller.go`, `service.go`, `repository.go`, `dto/`, and `index.go` (its `Register*Routes`).
- Route registration: public routes go before `api.Use(middleware.JWTProtected())` in `router/router.go`, protected routes after. New modules must be wired into `SetupRoutes`.
- Response/error helpers, JWT, pagination, and validation live in `utils/` — use them instead of hand-rolling responses.

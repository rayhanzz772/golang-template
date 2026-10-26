.PHONY: run build seed migrate help

BINARY=bin/app

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "  run      Start the HTTP server"
	@echo "  build    Build the CLI binary to $(BINARY)"
	@echo "  seed     Seed the database"
	@echo "  migrate  Run database auto-migration"

run:
	go run . server

build:
	go build -o $(BINARY) .

seed:
	go run . seed

migrate:
	go run . migrate

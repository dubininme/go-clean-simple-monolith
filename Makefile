.PHONY: help dev prod build up down logs migrate migrate-up migrate-down lint go-sdk

help:
	@echo "Project Makefile commands:"
	@echo "  dev          - Run dev environment (with hot-reload, dlv, Dev.Dockerfile)"
	@echo "  prod         - Run prod environment (production Dockerfile)"
	@echo "  build        - Build all containers (dev)"
	@echo "  up           - Start docker-compose (dev)"
	@echo "  down         - Stop and remove containers"
	@echo "  logs         - Show logs for all services"
	@echo "  migrate-up   - Apply all up migrations (docker-compose run migrate)"
	@echo "  migrate-down - Rollback last migration (docker-compose run migrate with down)"
	@echo "  lint         - Run golangci-lint"
	@echo "  go-sdk       - Generate Go SDK from OpenAPI specification"

# Variables for docker-compose
PLATFORM ?= amd64
DOCKERFILE ?= Dev.Dockerfile
GITHUB_USER ?= your_user
GITHUB_TOKEN ?= your_token
OPENAPI_FILE=internal/delivery/http/v1/openapi.yaml

# Dev mode (with hot-reload, dlv)
dev:
	PLATFORM=$(PLATFORM) GITHUB_USER=$(GITHUB_USER) GITHUB_TOKEN=$(GITHUB_TOKEN) \
		docker-compose up -d --build

# Build containers (dev)
build:
	PLATFORM=$(PLATFORM) GITHUB_USER=$(GITHUB_USER) GITHUB_TOKEN=$(GITHUB_TOKEN) \
		docker-compose build

# Start docker-compose (dev)
run:
	PLATFORM=$(PLATFORM) GITHUB_USER=$(GITHUB_USER) GITHUB_TOKEN=$(GITHUB_TOKEN) \
		docker-compose up -d

# Stop and remove containers
down:
	docker-compose down

# Show logs for all services
logs:
	docker-compose logs -f --tail=100

# Apply all up migrations
migrate-up:
	docker-compose run --rm migrate

# Rollback last migration
migrate-down:
	docker-compose run --rm migrate migrate -path=/migrations -database "mysql://dev_user:dev_password@tcp(db:3306)/test_database" down 1

# Run golangci-lint
lint:
	golangci-lint run --timeout=2m ./...

# Generate Go SDK from OpenAPI specification
go-sdk:
	docker compose exec app mkdir -p pkg/gen/oapi
	docker compose exec app oapi-codegen -generate=types,client -package=oapi /go/src/app/${OPENAPI_FILE} > pkg/gen/oapi/oapi.go

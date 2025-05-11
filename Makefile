.PHONY: help dev prod build up down logs migrate migrate-up migrate-down

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

# Variables for docker-compose
PLATFORM ?= amd64
DOCKERFILE ?= Dev.Dockerfile
GITHUB_USER ?= your_user
GITHUB_TOKEN ?= your_token

# Dev mode (with hot-reload, dlv)
dev:
	PLATFORM=$(PLATFORM) GITHUB_USER=$(GITHUB_USER) GITHUB_TOKEN=$(GITHUB_TOKEN) \
		docker-compose up --build

# Production mode (production Dockerfile)
prod:
	docker-compose -f docker-compose.yml up --build

# Build containers (dev)
build:
	PLATFORM=$(PLATFORM) GITHUB_USER=$(GITHUB_USER) GITHUB_TOKEN=$(GITHUB_TOKEN) \
		docker-compose build

# Start docker-compose (dev)
up:
	PLATFORM=$(PLATFORM) GITHUB_USER=$(GITHUB_USER) GITHUB_TOKEN=$(GITHUB_TOKEN) \
		docker-compose up

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

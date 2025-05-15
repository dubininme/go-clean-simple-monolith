# Go Clean Monolith Template

This project is a **template for building scalable, maintainable Go monolith applications**. It demonstrates a clean, modular architecture inspired by Domain-Driven Design (DDD) and includes a working example of an **Ecommerce** domain. The template is designed for teams who want a single codebase with clear separation of concerns, supporting both HTTP API and asynchronous event processing as separate deployable services.

---

## Features

- **Monolithic codebase, modular deployment:**  
  Shared domain, application, and infrastructure code, but separate builds for:
  - `api`: HTTP API service
  - `worker`: Background/async event processor

- **Domain-Driven Design patterns:**  
  - Aggregates, Entities, Value Objects, Repositories, Domain Services, Application Services, Unit of Work, and Events.

- **Ecommerce Example:**  
  - Orders, Products, Shipments, Payments, and Outbox pattern for reliable event publishing.

- **Modern infrastructure:**  
  - MySQL, RabbitMQ, Elasticsearch (all via Docker Compose)
  - Hot-reload and remote debugging (Delve, Reflex)
  - OpenAPI-driven API (with Go SDK generation)

---

## Project Structure

```
cmd/                # Entrypoints for api and worker binaries
internal/
  domain/           # DDD domain: aggregates, entities, value objects, services, events, repositories
  application/      # Application layer: use cases, commands, queries, unit of work, event handlers
  infrastructure/   # Infrastructure: DB, broker, external clients, scheduler
  delivery/         # Delivery layer: HTTP handlers, middleware, broker consumers/publishers
db/migrations/      # SQL migrations
pkg/                # Shared packages (auth, logger, generated SDKs)
tests/              # Integration and fixtures
```

---

## Local Development

### Prerequisites

- [Docker](https://www.docker.com/)
- [Make](https://www.gnu.org/software/make/)
- (Optional) [GoLand](https://www.jetbrains.com/go/) or VSCode for debugging

### 1. Clone the repository

```sh
git clone <your-repo-url>
cd go-clean-monolith
```

### 2. Set up environment variables

Copy `.env.example` to `.env` and fill in any required values.

You will need a GitHub user/token for private Go module access:

```sh
export GITHUB_USER=your_user
export GITHUB_TOKEN=your_token
```

### 3. Build and start all services

```sh
make dev
```

This will:
- Build all containers
- Start MySQL, RabbitMQ, Elasticsearch, API, and Worker services
- Enable hot-reload for Go code (via Reflex)
- Expose API on [http://localhost:7112](http://localhost:7112)
- Expose Delve debugger on ports 8012 (api) and 8013 (worker)

#### For Apple Silicon (M1/M2):

```sh
PLATFORM=arm64 make dev
```

#### For x86_64 (default):

```sh
PLATFORM=amd64 make dev
```

### 4. Run database migrations

```sh
make migrate-up
```

To rollback the last migration:

```sh
make migrate-down
```

### 5. View logs

```sh
make logs
```

### 6. Lint the code

```sh
make lint
```

### 7. Generate Go SDK from OpenAPI

```sh
make go-sdk
```

---

## API Example

The Ecommerce API exposes endpoints for managing orders and products.  
See [`internal/delivery/http/v1/openapi.yaml`](internal/delivery/http/v1/openapi.yaml) for the full OpenAPI spec.

Example endpoints:
- `POST /orders` — Create a new order
- `GET /orders` — List all orders
- `GET /orders/{id}` — Get order by ID
- `POST /orders/{id}/paid` — Mark order as paid

---

## Extending the Template

- Add new domains by following the DDD structure in `internal/domain`
- Implement new use cases in `internal/application/usecase`
- Add new delivery mechanisms (e.g., gRPC, GraphQL) in `internal/delivery`
- Use the outbox pattern for reliable event-driven workflows

---
# products-api-go

REST API for product management built with Go and hexagonal architecture.

## Tech Stack

- **Go** — main language
- **Gin** — HTTP framework
- **PostgreSQL** — database
- **pgx** — PostgreSQL driver

## Architecture

This project follows **hexagonal architecture** (Ports & Adapters), keeping the business logic decoupled from external dependencies.

```
internal/
├── domain/         # Core entities and business rules
├── ports/
│   ├── input/      # Interfaces exposed to the outside (use cases)
│   └── output/     # Interfaces for external dependencies (DB, etc.)
├── application/    # Use case implementations
└── adapters/
    ├── http/       # HTTP handlers and router (Gin)
    └── postgres/   # PostgreSQL repository implementation
```

## Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | List all products |
| GET | `/products/:id` | Get a product by ID |
| POST | `/products` | Create a product |
| PUT | `/products/:id` | Update a product |
| DELETE | `/products/:id` | Delete a product |

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 18+

### Database setup

```sql
CREATE DATABASE productos_db;

\c productos_db

CREATE TABLE products (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    price       NUMERIC(10,2) NOT NULL,
    stock       INTEGER NOT NULL DEFAULT 0
);
```

### Environment variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `` | Database password |
| `DB_NAME` | `productos_db` | Database name |
| `SERVER_PORT` | `8080` | HTTP server port |

### Run

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## Example Request

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Chair","description":"Office chair","price":99.99,"stock":10}'
```

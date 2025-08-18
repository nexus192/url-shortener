# URL Shortener

A clean and modular URL shortener service written in Go, following clean architecture principles.

## Project Structure

## Features

- Shorten long URLs into unique IDs.
- Resolve short IDs back to original URLs.
- Persistent storage in PostgreSQL.
- Separation of concerns: clean architecture with domain, data, and adapter layers.
- Easy to extend or replace storage layer.

### Prerequisites

- Go 1.21+
- PostgreSQL
- `pgx` driver for PostgreSQL

### Configuration

Create a `config/config.yaml` file with the following structure:

```yaml
database_url: "postgres://user:password@localhost:5432/url_shortener?sslmode=disable"
server_port: 8080

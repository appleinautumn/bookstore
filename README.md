# Book Store service

A simple service that serves Book Store API.

## Requirements

This project is developed with:

- Go 1.22
- Postgres 16

## Installation

Clone the project

```bash
git clone git@github.com:appleinautumn/bookstore.git
```

Go to the project directory

```bash
cd bookstore
```

This service contains a `.env.example` file that defines environment variables you need to set. Copy and set the variables to a new `.env` file.

```bash
cp .env.example .env
```

Start the app

```bash
go run cmd/api/api.go
```

## Database

If you have not created the database for Book Store service, please create one before going to the next step.

We're using [golang-migrate](https://github.com/golang-migrate/migrate) for the migration.

Install the package

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Run the migration (change the value accordingly)

```bash
$ migrate -path internal/migrations -database "postgres://postgres:password@127.0.0.1:5432/database?sslmode=disable&search_path=public" up
1/u create_books_table (13.118875ms)
2/u create_users_table (27.830055ms)
3/u create_orders_table (38.646938ms)
4/u create_order_books_table (50.738179ms)

```

To rollback

```bash
$ migrate -path internal/migrations -database "postgres://postgres:password@127.0.0.1:5432/database?sslmode=disable&search_path=public" down 1
1/d create_books_table (41.936181ms)
```

### Running Seeds

Seeding exists as a separate app since golang-migrate doesn't have seeding feature.
Run the following command:

```bash
$ go run cmd/seed/seed.go
table books seeded
```

## Testing

Run testing with coverage

```bash
go test -coverprofile=coverage.out ./...
```

Show coverage detail

```bash
go tool cover -func=coverage.out
```

Show coverage detail as HTML

```bash
go tool cover -html=coverage.out
```

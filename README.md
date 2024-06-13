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
4/d create_order_books_table (18.151815ms)
```

### Running Seeds

Seeding exists as a separate app since golang-migrate doesn't have seeding feature.
Run the following command:

```bash
$ go run cmd/seed/seed.go
table books seeded
```

## API Routes

### Public Routes

#### List Books

<details>
  <summary><code>POST</code> <code><b>/books</b></code></summary>

##### Request Body

> None

##### Responses

> | http code | content-type                      | response                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
> | --------- | --------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
> | `200`     | `application/json; charset=UTF-8` | `{"data":[{"id":1,"title":"The Great Gatsby","author":{"String":"F. Scott Fitzgerald","Valid":true},"description":{"String":"The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald. Set in the Jazz Age on Long Island, near New York City, the novel depicts first-person narrator Nick Carraway's interactions with mysterious millionaire Jay Gatsby and Gatsby's obsession to reunite with his former lover, Daisy Buchanan.","Valid":true},"created_at":"2024-06-12T14:04:01.875488Z","updated_at":"2024-06-12T14:04:01.875488Z"},{"id":2,"title":"To Kill a Mockingbird","author":{"String":"Harper Lee","Valid":true},"description":{"String":"To Kill a Mockingbird is a novel by Harper Lee published in 1960. Instantly successful, widely read in high schools and middle schools in the United States, it has become a classic of modern American literature, winning the Pulitzer Prize.","Valid":true},"created_at":"2024-06-12T14:04:01.877057Z","updated_at":"2024-06-12T14:04:01.877057Z"}]}` |

##### Example cURL

> ```javascript
>  curl --location 'http://localhost:8080/books'
> ```

</details>

#### Sign Up

<details>
  <summary><code>POST</code> <code><b>/signup</b></code></summary>

##### Request Body

> | name     | type     | data type | description |
> | -------- | -------- | --------- | ----------- |
> | Email    | required | string    | N/A         |
> | Name     | required | string    | N/A         |
> | Password | required | string    | N/A         |

##### Responses

> | http code | content-type                      | response                                                                                                                                                         |
> | --------- | --------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
> | `201`     | `application/json; charset=UTF-8` | `{"data":{"id":8,"email":"h2h@test.com","name":"fg","password":"passss","created_at":"2024-06-12T14:19:15.792706Z","updated_at":"2024-06-12T14:19:15.792706Z"}}` |
> | `400`     | `application/json; charset=UTF-8` | `{"error":{"message":"`email`must be`required=`"}}`                                                                                                              |
> | `500`     | `application/json; charset=UTF-8` | `{"error": {"message": "pq: duplicate key value violates unique constraint \"users_email_key\""}}`                                                               |

##### Example cURL

> ```javascript
> curl --location 'http://localhost:8080/signup' \
> --header 'Content-Type: application/json' \
> --data-raw '{
>     "email":"h2h@test.com",
>     "name":"fg",
>     "password":"passss"
> }'
> ```

</details>

### Private Routes

#### Order

<details>
  <summary><code>POST</code> <code><b>/my/orders</b></code></summary>

##### Headers

> | key     | value    | data type | description |
> | ------- | -------- | --------- | ----------- |
> | user_id | required | string    | N/A         |

##### Request Body

> | name   | type     | data type | description |
> | ------ | -------- | --------- | ----------- |
> | orders | required | array     | N/A         |

##### Responses

> | http code | content-type                      | response                                                                                                                   |
> | --------- | --------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
> | `200`     | `application/json; charset=UTF-8` | `{"data":{"id":7,"user_id":6,"created_at":"2024-06-12T14:26:45.095904Z","updated_at":"2024-06-12T14:26:45.095904Z"}}`      |
> | `400`     | `application/json; charset=UTF-8` | `{"error":{"message":"EOF"}}`                                                                                              |
> | `500`     | `application/json; charset=UTF-8` | `{"error":{"message":"pq: insert or update on table \"orders\" violates foreign key constraint \"orders_user_id_fkey\""}}` |

##### Example cURL

> ```javascript
> curl --location 'http://localhost:8080/orders' \
> --header 'user_id: 6' \
> --header 'Content-Type: application/json' \
> --data '{
>   "orders": [
>     {
>       "book_id": 1,
>       "quantity": 2
>     },
>     {
>       "book_id": 5,
>       "quantity": 1
>     }
>   ]
> }'
> ```

</details>

#### List Orders

<details>
  <summary><code>GET</code> <code><b>/my/orders</b></code></summary>

##### Headers

> | key     | value    | data type | description |
> | ------- | -------- | --------- | ----------- |
> | user_id | required | string    | N/A         |

##### Request Body

> None

##### Responses

> | http code | content-type                      | response                                                                                                                |
> | --------- | --------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
> | `200`     | `application/json; charset=UTF-8` | `{"data":[{"id":5,"user_id":5,"created_at":"2024-06-12T14:04:01.887702Z","updated_at":"2024-06-12T14:04:01.887702Z"}]}` |

##### Example cURL

> ```javascript
> curl --location 'http://localhost:8080/orders' \
> --header 'user_id: 5'
> ```

</details>

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

# A simple User API:

```
POST   /users
GET    /users
GET    /users/{id}
PUT    /users/{id}
DELETE /users/{id}`
```

# Foldert structure

```
go-user-api/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── db/
│   ├── models/
│   ├── repository/
│   ├── service/
│   └── handler/
│
├── migrations/
├── docker/
│
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```

# Comparison to spring boot

```
•	handler ≈ Controller
•	service ≈ Service
•	repository ≈ Repository
•	models ≈ Entities
```

# Dependencies

```
github.com/gin-gonic/gin - HTTP web framework
github.com/jackc/pgx/v5 - PostgreSQL driver + toolkit
github.com/jackc/pgx/v5/pgxpool - connection pool manager
github.com/joho/godotenv - loads environment variables from a .env file
```

# Testing locally

- docker-compose up --build
- docker exec - it postgres-go psql -U postgres -d users_db
- create a users table inside a postgres-go container:
    ```
     CREATE TABLE users (
         id SERIAL PRIMARY KEY,
         name TEXT NOT NULL,
         email TEXT UNIQUE NOT NULL
     );
  ```
- create a user using the API:
  ```
  curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Luiz","email":"luiz@email.com"}'
  ```
- retrieving all users using the API:

  ```curl -X GET http://localhost:8080/users```

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
go get github.com/gin-gonic/gin - HTTP web framework
go get github.com/jackc/pgx/v5 - PostgreSQL driver + toolkit
go get github.com/jackc/pgx/v5/pgxpool - connection pool manager
go get github.com/joho/godotenv - loads environment variables from a .env file
```

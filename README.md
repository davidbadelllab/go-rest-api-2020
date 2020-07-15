# Go REST API - Product Service

RESTful microservice built with Go 1.15, Gin framework, and PostgreSQL.

## Features

- ✅ REST API with Gin Web Framework
- ✅ PostgreSQL database with GORM
- ✅ JWT authentication
- ✅ Middleware (CORS, Auth, Logging)
- ✅ Input validation
- ✅ Dockerized application
- ✅ Swagger documentation

## Technologies

- **Go 1.15**
- **Gin Web Framework**
- **GORM** (PostgreSQL ORM)
- **JWT-Go** for authentication
- **Swagger** for API docs
- **Docker & Docker Compose**

## Go 1.15 Features

- Small allocation improvements
- New embedded tzdata package
- time/tzdata package

## Installation

```bash
# Clone repository
git clone https://github.com/davidbadelllab/go-rest-api-2020

# Install dependencies
go mod download

# Run with Docker
docker-compose up --build

# Or run locally
go run main.go
```

## API Endpoints

### Auth
```
POST /api/auth/register
POST /api/auth/login
```

### Products (Protected)
```
GET    /api/products
GET    /api/products/:id
POST   /api/products
PUT    /api/products/:id
DELETE /api/products/:id
```

## Example Request

```bash
# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret123",
    "name": "John Doe"
  }'

# Get products
curl -H "Authorization: Bearer <token>" \
  http://localhost:8080/api/products
```

## Project Structure

```
go-rest-api-2020/
├── main.go
├── config/
│   └── database.go
├── controllers/
│   ├── auth_controller.go
│   └── product_controller.go
├── models/
│   ├── user.go
│   └── product.go
├── middleware/
│   ├── auth.go
│   └── cors.go
├── routes/
│   └── routes.go
└── utils/
    └── jwt.go
```

## Environment Variables

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_rest_api
JWT_SECRET=your-secret-key
```

## Testing

```bash
go test ./...
```

## Author

David Badell - 2020

## License

MIT

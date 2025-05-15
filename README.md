# Gin Starter

<p align="center">
  <img src="https://img.shields.io/badge/Gin-1.9.1-green?logo=go" alt="Gin" />
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go" alt="Go" />
  <img src="https://img.shields.io/badge/GORM-2.0+-blue?logo=go" alt="GORM" />
  <img src="https://img.shields.io/badge/PostgreSQL-14+-336791?logo=postgresql" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Redis-7+-DC382D?logo=redis" alt="Redis" />
  <img src="https://img.shields.io/badge/Docker-20+-2496ED?logo=docker" alt="Docker" />
  <img src="https://img.shields.io/badge/Swagger-UI-85EA2D?logo=swagger" alt="Swagger" />
</p>

## 🚀 Introduction

**Gin Starter** is a backend boilerplate project built with [Gin](https://github.com/gin-gonic/gin) (Go). It is designed to help you quickly start a new backend service with a clean architecture and common integrations already set up. The template includes:

- Database integration with PostgreSQL using GORM ORM.
- Redis caching support.
- JWT-based authentication.
- Swagger UI for API documentation.
- Docker and docker-compose for containerized development and deployment.

This project follows a modular structure, separating concerns into application, domain, infrastructure, and interface layers, making it easy to extend and maintain.

## 🛠️ Technologies Used

| Technology     | Description                  | Icon |
| -------------- | --------------------------- | ---- |
| Go             | Programming language         | ![Go](https://img.shields.io/badge/Go-00ADD8?logo=go) |
| Gin            | Web framework                | ![Gin](https://img.shields.io/badge/Gin-1.9.1-green?logo=go) |
| GORM           | ORM for Go                   | ![GORM](https://img.shields.io/badge/GORM-2.0+-blue?logo=go) |
| PostgreSQL     | Relational database          | ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?logo=postgresql) |
| Redis          | Caching                      | ![Redis](https://img.shields.io/badge/Redis-DC382D?logo=redis) |
| Docker         | Packaging & Deployment       | ![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker) |
| Swagger        | API Documentation            | ![Swagger](https://img.shields.io/badge/Swagger-UI-85EA2D?logo=swagger) |

## 📦 Project Structure

```
gin-starter/
├── build/package/Dockerfile
├── cmd/server/
├── deployments/docker-compose.yml
├── docs/
├── internal/
│   ├── application/
│   ├── domain/
│   ├── infrastructure/
│   ├── interface/
│   └── shared/
└── ...
```

- **build/package/Dockerfile**: Dockerfile for building the Go application image.
- **cmd/server/**: Entry point for the application (main.go).
- **deployments/docker-compose.yml**: Docker Compose file for local development.
- **docs/**: Swagger/OpenAPI documentation files.
- **internal/application/**: Application logic, DTOs, and services.
- **internal/domain/**: Domain entities and repository interfaces.
- **internal/infrastructure/**: Database, caching, and repository implementations.
- **internal/interface/**: HTTP handlers, routers, and middleware.
- **internal/shared/**: Shared constants, utilities, and dependency injection.

## ⚡️ Quick Start

### 1. Clone the project

```bash
git clone https://github.com/your-username/gin-starter.git
cd gin-starter
```

### 2. Create `.env` file

Create a `.env` file in the root directory with the following content (adjust values as needed):

```env
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=gin_starter
POSTGRES_SSL_MODE=disable
POSTGRES_TIME_ZONE=Asia/Ho_Chi_Minh

JWT_SECRET_KEY=your_jwt_secret

REDIS_CACHING_HOST=localhost
REDIS_CACHING_PORT=6379
REDIS_CACHING_PASSWORD=
REDIS_CACHING_DB=0
```

### 3. Run with Docker Compose

To start the database, Redis, and (optionally) the backend service in containers:

```bash
docker compose -f deployments/docker-compose.yml up --build
```

> **Note:** To run the backend in Docker, uncomment the `gin-starter` service in `docker-compose.yml`.

### 4. Run locally (without Docker)

If you want to run the backend directly on your machine (make sure PostgreSQL and Redis are running):

```bash
go run cmd/server/main.go
```

### 5. Access Swagger UI

After starting the backend, you can view the API documentation at:

- [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

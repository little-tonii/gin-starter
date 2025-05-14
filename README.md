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

**Gin Starter** is a backend project template using [Gin](https://github.com/gin-gonic/gin) (Go), pre-integrated with popular components such as:

- **GORM** for ORM with PostgreSQL
- **Redis** for caching
- **JWT** for authentication
- **Swagger** for API documentation
- **Docker** & **docker-compose** for development and deployment

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

## ⚡️ Quick Start

### 1. Clone the project

```bash
git clone https://github.com/your-username/gin-starter.git
cd gin-starter
```

### 2. Create `.env` file

Create a `.env` file in the root directory with content like:

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

```bash
docker compose -f deployments/docker-compose.yml up --build
```

> **Note:** To run the backend with Docker, uncomment the `gin-starter` service in `docker-compose.yml`.

### 4. Run locally (without Docker)

```bash
go run cmd/server/main.go
```

### 5. Access Swagger UI

- [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
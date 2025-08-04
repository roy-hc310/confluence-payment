## 🚀 Getting Started

### Prerequisites

- Go 1.23.4 or higher
- Docker and Docker Compose

### Installation

1. Create directory:
```bash
mkdir /confluence
git clone https://github.com/yourusername/confluence-payment.git
```

2. Add this docker-compose.yaml to /confluence directory:
```bash
version: '3.8'

services:
  confluence-order:
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 500M
        reservations:
          cpus: '0.5'
          memory: 500M
    build:
      context: ./confluence-order
      dockerfile: Dockerfile
    ports:
      - 8000:8000
    container_name: confluence-order
    depends_on:
      postgres:
        condition: service_healthy
      redis: 
        condition: service_healthy

  confluence-payment:
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 500M
        reservations:
          cpus: '0.5'
          memory: 500M
    build:
      context: ./confluence-payment
      dockerfile: Dockerfile
    ports:
      - 8001:8001
      - 50052:50052
    container_name: confluence-payment
    depends_on:
      postgres:
        condition: service_healthy
      redis: 
        condition: service_healthy


  postgres:
    image: postgres:15
    container_name: confluence-postgres
    restart: always
    environment:
      - POSTGRES_DB=confluence
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=admin
    ports:
      - '5433:5432'
    volumes:
      - postgres_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres -d confluence"]
      interval: 10s
      timeout: 5s
      retries: 5
      start_period: 10s

  redis:
    image: redis:7
    container_name: confluence-redis
    restart: always
    ports:
      - '6379:6379'
    volumes:
      - redis_data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 5s
      retries: 5
    
volumes:
  postgres_data:
  redis_data:
```

3. Setup database:
```bash
# Go to Postgres port 5433 and database named "confluence"
CREATE DATABASE confluence;

# Create 2 schemas in database confluence
CREATE SCHEMA IF NOT EXISTS "order"; 
CREATE SCHEMA IF NOT EXISTS "payment";

# Copy queries from confluence/confluence-payment/core-internal/migrations/000001_init_tables.up.sql and run the queries to create table
```

4. API collection:
```bash
# go to postman pulic collection 
# https://www.postman.com/dark-eclipse-55522/workspace/spgroup/request/20536686-16ee6ffa-a068-4c9e-a921-a2a1efa5cd76?action=share&creator=20536686&ctx=documentation
```



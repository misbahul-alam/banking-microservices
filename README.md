# Banking Microservices

A production-grade banking backend built with **Go** using a **microservices architecture**. The project demonstrates modern backend concepts such as API Gateway, gRPC, Kafka, PostgreSQL, Redis, and Docker.

## Architecture

```text
Client
   │
   ▼
API Gateway (REST)
   │
   ├──────────────┬──────────────┐
   ▼              ▼              ▼
 Auth         Account      Transaction
   ▲              ▲              │
   └─────── gRPC ─┴──────────────┘
                  │
                  ▼
                Kafka
                  │
                  ▼
           Notification
```

## Tech Stack

- Go
- Chi
- gRPC
- Apache Kafka
- PostgreSQL
- Redis
- SQLC + pgx
- JWT
- Docker & Docker Compose

## Project Structure

```text
banking-microservices/
├── services/
│   ├── api-gateway/
│   ├── auth/
│   ├── account/
│   ├── transaction/
│   └── notification/
├── proto/
├── deployments/
├── docs/
├── scripts/
├── docker-compose.yml
├── go.work
└── README.md
```

## Services

- **API Gateway** – Routes client requests
- **Auth** – Authentication & JWT
- **Account** – Account management
- **Transaction** – Deposit, Withdraw, Transfer
- **Notification** – Consumes Kafka events

## Communication

- **REST** → Client → API Gateway
- **gRPC** → Service-to-Service
- **Kafka** → Event Streaming

## Getting Started

```bash
git clone https://github.com/misbahul-alam/banking-microservices.git
cd banking-microservices

go work sync

docker compose up -d
```

## Roadmap

- [ ] API Gateway
- [ ] Authentication
- [ ] Account Management
- [ ] Transactions
- [ ] Kafka Integration
- [ ] Redis
- [ ] Docker
- [ ] CI/CD
- [ ] Kubernetes

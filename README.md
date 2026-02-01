# Marketplace Engine

A distributed marketplace backend built with Go. This project demonstrates a microservice architecture featuring isolated databases, automated schema migrations, and inter-service synchronization via gRPC.

## Architectural Design

The system is built with strict service boundaries:

- **Core API Service**: Manages the product catalog, user authentication, and transactional order processing.
- **Partner API Service**: B2B gateway for suppliers. It manages partner-specific data and uses gRPC to synchronize inventory with the Core API.
- **Isolated Persistence**: Each service owns its own PostgreSQL instance.

### Key Technical Features
* **Schema-First Development**: API contracts defined using OpenAPI 3.0 (REST) and Protocol Buffers (gRPC).
* **Inter-Service Communication**: Partner API communicates with Core API via gRPC for low-latency.
* **Automated Migrations**: Integrated Flyway migration engine.
* **Infrastructure as Code**: Full system orchestration via Docker Compose with healthchecks and dependency management.

---

## Tech Stack

- **Backend**: Go (Golang)
- **Communication**: gRPC, Protobuf, REST (OpenAPI 3.0)
- **Databases**: PostgreSQL
- **Migrations**: Flyway
- **Orchestration**: Docker & Docker Compose

---

## Networking

The system is fully containerized, utilizing Docker's internal DNS for service discovery.



| Service | Internal Port | External Port | Role |
| :--- | :--- | :--- | :--- |
| `api` | `8000` (HTTP) | `8000` | Customer Facing API |
| `api` | `50052` (gRPC) | - | Internal Inventory RPC |
| `partner_api` | `8001` (HTTP) | `8001` | B2B Partner Gateway |
| `db_marketplace` | `5432` | `5434` | Storage for Core API |
| `db_partners` | `5432` | `5433` | Storage for Partner API |

---

## Getting Started

### Prerequisites
* Docker and Docker Compose installed.

### Installation & Launch
The entire stack, including databases and migrations, starts with one command:

```bash

docker-compose up --build
```

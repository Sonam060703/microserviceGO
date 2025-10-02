
# MicroserviceGO

A scalable e-commerce microservices architecture built with Go, featuring gRPC inter-service communication, GraphQL API gateway, and containerized deployment.

## 🏗️ Architecture Overview

```
                                    ┌─────────────────┐
                                    │   GraphQL API   │
                                    │    Gateway      │
                                    │   (Port 8000)   │
                                    └────────┬────────┘
                                             │
                    ┌────────────────────────┼────────────────────────┐
                    │                        │                        │
            ┌───────▼────────┐      ┌───────▼────────┐      ┌───────▼────────┐
            │    Account     │      │    Catalog     │      │     Order      │
            │    Service     │      │    Service     │      │    Service     │
            │   (gRPC:8080)  │      │   (gRPC:8080)  │      │   (gRPC:8080)  │
            └───────┬────────┘      └───────┬────────┘      └───────┬────────┘
                    │                        │                        │
            ┌───────▼────────┐      ┌───────▼────────┐      ┌───────▼────────┐
            │   PostgreSQL   │      │ Elasticsearch  │      │   PostgreSQL   │
            │  (Account DB)  │      │   + Kibana     │      │   (Order DB)   │
            └────────────────┘      └────────────────┘      └────────────────┘
```

## 📋 Services

### 1. **Account Service**
- Manages user accounts and authentication
- gRPC-based service
- PostgreSQL database for data persistence
- Endpoints for user registration, login, and profile management

### 2. **Catalog Service**
- Handles product catalog management
- Elasticsearch integration for fast product search
- Product CRUD operations
- Advanced search and filtering capabilities

### 3. **Order Service**
- Manages customer orders and order processing
- Communicates with Account and Catalog services
- PostgreSQL database for order persistence
- Order lifecycle management (create, update, track)

### 4. **GraphQL Gateway**
- Unified API entry point for client applications
- Aggregates data from multiple microservices
- Exposed on port 8000
- Provides flexible querying capabilities

## 🚀 Technology Stack

- **Language**: Go (Golang)
- **Communication**: gRPC for inter-service communication
- **API Gateway**: GraphQL
- **Databases**: 
  - PostgreSQL (Account & Order services)
  - Elasticsearch (Catalog service)
- **Containerization**: Docker & Docker Compose
- **Monitoring**: Kibana for Elasticsearch visualization

## 📦 Project Structure

```
microserviceGO/
├── account/              # Account microservice
│   ├── cmd/             # Service entry point
│   ├── pb/              # Protocol buffer definitions
│   ├── account.proto    # gRPC service definition
│   ├── repository.go    # Database layer
│   ├── service.go       # Business logic
│   └── server.go        # gRPC server
├── catalog/             # Catalog microservice
│   ├── cmd/            # Service entry point
│   ├── pb/             # Protocol buffer definitions
│   ├── catalog.proto   # gRPC service definition
│   ├── repository.go   # Elasticsearch integration
│   ├── service.go      # Business logic
│   └── server.go       # gRPC server
├── order/              # Order microservice
│   ├── cmd/           # Service entry point
│   ├── pb/            # Protocol buffer definitions
│   ├── order.proto    # gRPC service definition
│   ├── repository.go  # Database layer
│   ├── service.go     # Business logic
│   └── server.go      # gRPC server
├── graphql/           # GraphQL gateway
│   ├── schema.graphql # GraphQL schema
│   ├── resolvers/     # Query and mutation resolvers
│   └── main.go        # Gateway entry point
└── docker-compose.yaml # Container orchestration
```

## 🛠️ Prerequisites

- Docker (version 20.x or higher)
- Docker Compose (version 2.x or higher)
- Go 1.21+ (for local development)

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Sonam060703/microserviceGO.git
cd microserviceGO
```

### 2. Start All Services

```bash
docker-compose up --build
```

This command will:
- Build all microservice images
- Start PostgreSQL databases for Account and Order services
- Start Elasticsearch and Kibana for the Catalog service
- Launch all microservices
- Expose the GraphQL gateway on port 8000

### 3. Access the Services

- **GraphQL Playground**: http://localhost:8000
- **Kibana Dashboard**: http://localhost:5601
- **Elasticsearch**: http://localhost:9200

## 🧪 Testing the API

### GraphQL Queries Example

```graphql
# Create a new account
mutation {
  createAccount(name: "John Doe", email: "john@example.com") {
    id
    name
    email
  }
}

# Query products
query {
  products(query: "laptop", skip: 0, take: 10) {
    id
    name
    description
    price
  }
}

# Create an order
mutation {
  createOrder(accountId: "user-id", products: [{id: "product-id", quantity: 2}]) {
    id
    totalPrice
    createdAt
  }
}
```

## 🔧 Configuration

Each service can be configured via environment variables in `docker-compose.yaml`:

- **Database URLs**: PostgreSQL connection strings
- **Service URLs**: Inter-service communication endpoints
- **Elasticsearch**: Connection to the search cluster

## 📊 Health Checks

All services include health checks:
- PostgreSQL databases: `pg_isready` checks
- Elasticsearch: Cluster health API endpoint
- Services automatically restart on failure

## 🏗️ Development

### Build Individual Services

```bash
# Build account service
cd account
go build -o bin/account ./cmd/account

# Build catalog service
cd catalog
go build -o bin/catalog ./cmd/catalog

# Build order service
cd order
go build -o bin/order ./cmd/order
```

### Generate Protocol Buffers

```bash
protoc --go_out=. --go-grpc_out=. account/account.proto
protoc --go_out=. --go-grpc_out=. catalog/catalog.proto
protoc --go_out=. --go-grpc_out=. order/order.proto
```

## 🔐 Security Considerations

- Update default database credentials in production
- Implement proper authentication and authorization
- Enable SSL/TLS for inter-service communication
- Use secrets management for sensitive data

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is open source and available under the MIT License.

## 👥 Author

**Sonam060703**
- GitHub: [@Sonam060703](https://github.com/Sonam060703)

## 🙏 Acknowledgments

- Built with Go and gRPC for high-performance microservices
- GraphQL for flexible API queries
- Docker for containerization and easy deployment
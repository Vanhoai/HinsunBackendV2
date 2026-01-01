# Hinsun API

A robust backend API for a portfolio and blog system built with Go, following Domain-Driven Design (DDD) and Hexagonal Architecture principles.

## 🏗️ Architecture

This project follows **Hexagonal Architecture** with **Domain-Driven Design (DDD)** principles:

```
├── adapters/               # Adapters layer
│   ├── primary/            # Incoming adapters (HTTP handlers)
│   │   ├── v1/             # API version 1
│   │   └── v2/             # API version 2
│   ├── secondary/          # Outgoing adapters (repositories, external APIs)
│   │   ├── apis/           # External API clients
│   │   └── repositories/   # Implementations of domain repositories
│   └── shared/             # Shared adapter code
│       ├── databases/      # Database connections
│       ├── di/             # Dependency injection
│       ├── https/          # Models and utilities for HTTP
│       ├── middlewares/    # HTTP middlewares
│       └── models/         # Data models (used by implementations of repositories)
├── cmd/                    # Application entry points
│   └── main.go
├── configs/                # Configuration management
├── internal/               # Internal application code
│   ├── core/               # Core utilities
│   │   ├── events/         # Event bus
│   │   ├── failure/        # Error handling
│   │   ├── log/            # Logging
│   │   └── types/          # Common types
│   └── domain/             # Domain layer
│       ├── account/
│       ├── applications/   # Application services
│       ├── auth/
│       ├── blog/
│       ├── experience/
│       ├── notification/
│       ├── project/
│       ├── usecases/       # Use case definitions
│       └── values/         # Value objects
└── pkg/                    # Packages
    ├── firebase/
    ├── jwt/
    └── security/
```

## 🚀 Features

- **Hexagonal Architecture**: Separation of concerns with clear boundaries
- **Domain-Driven Design**: Rich domain models with business logic
- **RESTful API**: Following REST principles with versioning support
- **JWT Authentication**: Secure authentication with jwt and RSA/ECDSA signing
- **Event-Driven**: Asynchronous event bus for decoupled communication
- **Dependency Injection**: Using Uber FX for clean DI
- **Hot Reload**: Development with Air for instant feedback
- **Structured Logging**: Using Zap with rotation support
- **Database**: PostgreSQL with GORM

## 🛠️ Tech Stack

- **Language**: Go 1.25.5
- **Framework**: Chi Router v5
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT with RSA/ECDSA signing
- **Password Hashing**: Argon2id
- **Logging**: Uber Zap with Lumberjack rotation
- **DI Container**: Uber FX
- **Validation**: go-playground/validator
- **Configuration**: godotenv

## 📋 Prerequisites

- Go 1.25.5 or higher
- PostgreSQL 18+ (because I use UUIDv7 function which is supported from this version)
- Air (for hot reload during development)

## 🔧 Installation

1. Clone the repository:

```bash
git clone https://github.com/Vanhoai/HinsunBackendV2
cd HinsunBackendV2
```

2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables:

```bash
cp .env.example .env
```

After copying, edit the `.env` file to configure your database and other settings.

4. Run database migrations or create tables manually:

```go
// go to this file: adapters/shared/databases/postgres_database.go
// uncomment the following line:

// gormDB.AutoMigrate(&models.ExperienceModel{}, &models.AccountModel{})  <- uncomment this line
```

## 🚀 Running the Application

### Development (with hot reload):

```bash
air
```

### Production:

```bash
go build -o app cmd/main.go
./app
```

The server will start on the configured address (default: `:8080`)

## 🧪 Testing

This feature is coming soon 🫣 !!

## 🏗️ Hexagonal & Domain Driven Design

1. **Domain Layer**: Contains business logic and entities

   Domain is layer most important in this architecture, it contains all business logic and rules. Following DDD principles, the domain layer is organized into aggregates, entities, value objects, domain services, and application services.

   In this project, the domain layer is structured as follows:

   - **Entities**: Represent core business objects with unique identities
   - **Value Objects**: Immutable objects representing descriptive aspects of the domain
   - **Repositories**: Interfaces for data access and persistence
   - **Domain Services**: Encapsulate domain logic that doesn't fit within entities or value objects
   - **Application Services**: Coordinate use cases and interact with domain services

   💡 Note: normally, application services should receive commands/queries with DTOs format and return DTOs as well. However, in this project, for simplicity, I let application services receive and return domain entities directly.

2. **Adapter Layer**: Implements interfaces for external systems

   The adapter layer is responsible for communication between the domain layer and external systems such as databases, web frameworks, and third-party services. It contains both primary adapters (for incoming requests) and secondary adapters (for outgoing requests).

   Note: In repositories implementations, I use models (data models) to interact with the database instead of using domain entities directly. This approach helps to decouple the domain layer from the persistence layer and allows for easier mapping between domain entities and database records.

💪🏻 Bonus: I defined core modules in the internal/core directory, which can be reused across different projects. At here, I implemented some essential modules such as event bus, logging, error handling, and common types.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

This is the first time, I have used DDD and Hexagonal Architecture, so if you have any suggestions or improvements, please feel free to open an issue or submit a pull request 😤.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## 🙏 Acknowledgments

- Clean Architecture by Robert C. Martin
- Domain-Driven Design by Eric Evans
- Go community for excellent libraries and tools

# LLM Interface Backend API

A Go-based REST API backend for enabling chat-based AI assistant interactions. This project provides secure user authentication, chat session management, message handling, and LLM integration capabilities.

## 🏗️ Architecture

This project follows Clean Architecture principles and uses the C4 framework for documentation. The system is designed to be scalable, maintainable, and easily testable.

### Tech Stack
- **Language**: Go 1.21+
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL
- **Authentication**: JWT tokens
- **Migrations**: golang-migrate
- **Documentation**: C4 framework

## 📁 Project Structure

```
asstBackend/
├── cmd/server/           # Application entry point
├── internal/
│   ├── api/             # HTTP API setup and routing
│   ├── config/          # Configuration management
│   ├── database/        # Database connection and migrations
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # HTTP middleware (auth, CORS, etc.)
│   ├── models/          # Data models and DTOs
│   ├── repository/      # Data access layer
│   └── service/         # Business logic layer
├── migrations/          # Database migration files
├── docs/
│   └── c4/             # C4 framework documentation
├── .github/
│   ├── workflows/      # GitHub Actions CI/CD
│   └── copilot-instructions.md
├── go.mod
├── go.sum
├── .env.example
├── .gitignore
└── README.md
```

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- PostgreSQL 12+
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/jay13jay/asstBackend.git
cd asstBackend
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your PostgreSQL configuration
```

4. Ensure PostgreSQL is running:
```bash
# The application will automatically create the database if it doesn't exist
# Just ensure PostgreSQL server is running and accessible
```

5. Start the server (includes automatic database creation and migrations):
```bash
go run cmd/server/main.go
```
```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080` by default.

## 📚 API Documentation

### Authentication Endpoints

#### POST /api/v1/auth/register
Register a new user account.

**Request Body:**
```json
{
  "email": "user@example.com",
  "username": "username",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "username": "username",
    "created_at": "2025-11-08T...",
    "updated_at": "2025-11-08T..."
  },
  "access_token": "jwt_token",
  "refresh_token": "refresh_token"
}
```

#### POST /api/v1/auth/login
Authenticate an existing user.

### Chat Management Endpoints

All chat endpoints require authentication via Bearer token.

#### POST /api/v1/chats
Create a new chat conversation.

#### GET /api/v1/chats
List all chats for the authenticated user.

#### GET /api/v1/chats/{id}
Get a specific chat by ID.

#### PUT /api/v1/chats/{id}
Update a chat's title or description.

#### DELETE /api/v1/chats/{id}
Delete a chat and all its messages.

### Message Endpoints

#### POST /api/v1/chats/{id}/messages
Send a new message to a chat.

#### GET /api/v1/chats/{id}/messages
Retrieve all messages for a chat.

#### PUT /api/v1/chats/{id}/messages/{messageId}
Update an existing message.

#### DELETE /api/v1/chats/{id}/messages/{messageId}
Delete a message.

## 🏛️ C4 Documentation

This project uses the C4 framework for architectural documentation:

- **[L1 Context](docs/c4/L1-context.md)**: System landscape and external dependencies
  - 📊 [Diagram](docs/diagrams/asstBackend-l1-context.drawio) | [View Online](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l1-context.drawio)
- **[L2 Container](docs/c4/L2-container.md)**: High-level architecture
  - 📊 [Diagram](docs/diagrams/asstBackend-l2-containers.drawio) | [View Online](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l2-containers.drawio)
- **[L3 Component](docs/c4/L3-component.md)**: Detailed component breakdown
  - 📊 [Diagram](docs/diagrams/asstBackend-l3-components.drawio) | [View Online](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l3-components.drawio)
- **[L4 Code](docs/c4/L4-code.md)**: Implementation specifications
  - 📊 [Diagram](docs/diagrams/asstBackend-l4-code.drawio) | [View Online](https://app.diagrams.net/?lightbox=1#Uhttps://raw.githubusercontent.com/jay13jay/asstAPI/main/docs/diagrams/asstBackend-l4-code.drawio)

### Architecture Diagrams

Visual representations of the architecture are available in the [docs/diagrams](docs/diagrams/) directory. All diagrams are editable in [diagrams.net](https://app.diagrams.net) and can be regenerated from the C4 documentation.

### Node Traversal Algorithm

The documentation follows a Depth-First Search approach to systematically refine each level until implementation specifications are clear and deterministic.

## 🧪 Development

### Running Tests
```bash
go test ./...
```

### Running with Live Reload (using air)
```bash
# Install air if not already installed
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Database Operations

**Note**: The application automatically creates the database and runs migrations on startup, so manual database setup is usually not needed.

#### Create a new migration
```bash
migrate create -ext sql -dir migrations -seq migration_name
```

#### Manual migration commands (if needed)
```bash
# Apply migrations manually
migrate -path migrations -database "postgres://username:password@localhost:5432/asstbackend?sslmode=disable" up

# Rollback migrations manually
migrate -path migrations -database "postgres://username:password@localhost:5432/asstbackend?sslmode=disable" down
```

#### Database Auto-Creation
The application includes automatic database creation functionality:
- Checks if the target database exists
- Creates the database if it doesn't exist
- Runs all pending migrations automatically
- Enables required PostgreSQL extensions (uuid-ossp)

## 🛡️ Security Features

- JWT-based authentication
- Password hashing with bcrypt
- CORS middleware
- Input validation and sanitization
- Rate limiting (planned)
- SQL injection protection via parameterized queries

## 🎯 Future Enhancements

- [ ] LLM provider integrations (OpenAI, Anthropic)
- [ ] Real-time messaging via WebSockets
- [ ] File upload/attachment support
- [ ] Rate limiting implementation
- [ ] Comprehensive logging and monitoring
- [ ] Docker containerization
- [ ] Kubernetes deployment manifests
- [ ] API documentation with Swagger/OpenAPI
- [ ] Integration tests
- [ ] Performance benchmarking

## 📝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/new-feature`
3. Make your changes and add tests
4. Commit your changes: `git commit -am 'Add new feature'`
5. Push to the branch: `git push origin feature/new-feature`
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) file for details.

## 🤝 Support

For support, please create an issue in the GitHub repository or contact the development team.

---

**Note**: This project is in active development. API endpoints and features may change before the first stable release.
# L3 Component - LLM Interface Backend API

**Status**: DRAFT
**Last Updated**: November 8, 2025

This document breaks down the major components inside the API container and their responsibilities, interfaces, and interactions. It maps closely to the current codebase and is intended to be implementation-accurate.

## High-level components

1. Handlers (HTTP Layer)
   - Files: `internal/handlers/*.go`
   - Responsibility: Parse and validate HTTP requests, call corresponding service methods, format HTTP responses and status codes.
   - Key handlers:
     - `UserHandler` — register, login, refresh token, profile
     - `ChatHandler` — create, list, get, update, delete chats
     - `MessageHandler` — create, list, update, delete messages for a chat
   - Error strategy: handlers translate service errors to HTTP responses (400, 401, 404, 500)

2. Middleware
   - Files: `internal/middleware/*.go`
   - Responsibility: Cross-cutting concerns: CORS, authentication (JWT), request logging, (placeholder) rate limiting.
   - Notable behavior: `AuthMiddleware` extracts `userID` and `email` into Gin context for handlers.

3. Services (Business Logic)
   - Files: `internal/service/*.go`
   - Responsibility: Implement application rules and orchestration between repositories and external services (LLM providers in future).
   - Key services:
     - `UserService` — user creation, authentication, token generation
     - `ChatService` — chat lifecycle and ownership checks
     - `MessageService` — message creation, retrieval, update, deletion
   - Error strategy: return domain errors that handlers translate to HTTP responses

4. Repositories (Data Access)
   - Files: `internal/repository/*.go`
   - Responsibility: Encapsulate DB SQL queries. Parameterized queries to prevent SQL injection.
   - Key repos:
     - `UserRepository` — CRUD for users
     - `ChatRepository` — CRUD and list queries for chats
     - `MessageRepository` — CRUD and listing for messages

5. Models / DTOs
   - Files: `internal/models/models.go`
   - Responsibility: Define core data structures and validation annotations used by Gin binding (request DTOs, response shapes, JWT claims).
   - Key types: `User`, `Chat`, `Message`, `CreateUserRequest`, `LoginRequest`, `CreateChatRequest`, `CreateMessageRequest`, `AuthResponse`, `Claims`

6. Database & Migrations
   - Files: `internal/database/connection.go`, `migrations/*.sql`
   - Responsibility: DB connection lifecycle, ensure database exists (auto-create if needed), and apply migrations at startup.
   - Notes: Migrations enable `uuid-ossp` extension and create `users`, `chats`, `messages` tables with appropriate indexes.

7. API Server
   - Files: `internal/api/server.go`, `cmd/server/main.go`
   - Responsibility: Wire middleware, handlers, repositories, and services; start the Gin server.

## Component Interactions (sequence)

Example: Create message flow
1. HTTP POST `/api/v1/chats/{id}/messages` reaches `MessageHandler.CreateMessage`.
2. Middleware `AuthMiddleware` already populated `userID` in context.
3. Handler binds `CreateMessageRequest` and calls `MessageService.CreateMessage(chatID, req)`.
4. `MessageService` constructs a `models.Message` and calls `MessageRepository.Create`.
5. `MessageRepository` persists message and returns the stored message.
6. Handler returns HTTP 201 with the created message JSON.

## Public interfaces and function contracts

- Handler methods accept `*gin.Context` and return HTTP responses. Inputs validated using Gin `binding` tags.
- Service methods return domain objects and `error`. Error texts are intentionally simple; consider using typed errors for machine handling.
- Repository methods return concrete model pointers and `error`.

## Ownership and Authorization

- Chat operations are protected by ownership checks in `ChatService`.
- Message mutation endpoints rely on chat ownership logic in services to prevent unauthorized edits (ensure additional checks if you intend per-message ownership in future).

## Future Component Considerations

- LLM integration component (adapter) for making requests to OpenAI/Anthropic and storing responses. This should exist at the service layer as an interface so it can be swapped for different providers.
- Caching and rate-limiting middleware using Redis.
- Background worker queue for async LLM calls (e.g., for long-running tasks).

## Acceptance Criteria for L3 -> L4

- Each component above has an explicit API/contract in L4
- All request/response DTOs are documented
- DB schema is fully documented (fields, types, indexes)
- Authentication flow fully specified (token format, expiry)


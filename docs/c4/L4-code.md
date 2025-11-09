# L4 Code - Implementation Specifications and API Contracts

**Status**: DRAFT
**Last Updated**: November 8, 2025

This level contains concrete implementation details, API contracts, database schema, validation rules, and examples sufficient for a software engineer to implement or integrate with the service.

## API Overview (base path `/api/v1`)

Authentication endpoints (public):

1) POST /api/v1/auth/register
- Request (201 -> created on success)
  - Content-Type: application/json
  - Body (CreateUserRequest):
    - email: string (required, email format)
    - username: string (required, 3-50 chars)
    - password: string (required, min 8 chars)
- Responses
  - 201 Created: returns `AuthResponse` JSON `{ user, access_token, refresh_token }`
  - 400 Bad Request: invalid payload
  - 500 Internal Server Error: failure creating user

2) POST /api/v1/auth/login
- Request (200 -> ok on success)
  - Body (LoginRequest):
    - email: string (required)
    - password: string (required)
- Responses
  - 200 OK: `AuthResponse`
  - 401 Unauthorized: invalid credentials
  - 400 Bad Request: invalid payload

3) POST /api/v1/auth/refresh (not yet implemented)
- Expected: Accepts refresh token, returns new access/refresh tokens
- Responses: 200 OK (new tokens), 401 Unauthorized


Protected endpoints (require Authorization: Bearer <access_token>):

Chats
- POST /api/v1/chats
  - Body: `CreateChatRequest` { title: required, description: optional }
  - Responses: 201 Created -> Chat
- GET /api/v1/chats
  - Responses: 200 OK -> []Chat
- GET /api/v1/chats/{id}
  - Responses: 200 OK -> Chat, 404 Not Found
- PUT /api/v1/chats/{id}
  - Body: `UpdateChatRequest` { title?, description? }
  - Responses: 200 OK -> Chat
- DELETE /api/v1/chats/{id}
  - Responses: 204 No Content on success

Messages (scoped under a chat)
- POST /api/v1/chats/{id}/messages
  - Body: `CreateMessageRequest` { content: required, role: one of [user, assistant, system] }
  - Responses: 201 Created -> Message
- GET /api/v1/chats/{id}/messages
  - Responses: 200 OK -> []Message (ordered by created_at asc)
- PUT /api/v1/chats/{id}/messages/{messageId}
  - Body: `UpdateMessageRequest` { content: required }
  - Responses: 200 OK -> Message
- DELETE /api/v1/chats/{id}/messages/{messageId}
  - Responses: 204 No Content

## Data shapes (models)

User (JSON)
- id: uuid (string)
- email: string
- username: string
- created_at: timestamp
- updated_at: timestamp

Chat (JSON)
- id: uuid
- user_id: uuid
- title: string
- description: string (nullable)
- created_at: timestamp
- updated_at: timestamp

Message (JSON)
- id: uuid
- chat_id: uuid
- content: string
- role: string (user|assistant|system)
- created_at: timestamp
- updated_at: timestamp

AuthResponse
- user: User
- access_token: string (JWT)
- refresh_token: string (JWT)

## Validation rules
- Email: must be a valid email per Gin binding `email` tag
- Username: min=3, max=50
- Password: min=8 chars
- CreateChatRequest.Title: required, min=1, max=255
- CreateMessageRequest.Role: must be one of `user`, `assistant`, `system`

## HTTP Status mapping (recommended)
- 200 OK — success responses for GET/PUT
- 201 Created — successful creation (POST)
- 204 No Content — successful deletion
- 400 Bad Request — invalid payloads or params
- 401 Unauthorized — missing/invalid token
- 403 Forbidden — authenticated but not permitted (future)
- 404 Not Found — requested resource not present
- 500 Internal Server Error — unexpected server failures

## Database schema (summary)

Table: users
- id UUID PRIMARY KEY DEFAULT uuid_generate_v4()
- email VARCHAR(255) UNIQUE NOT NULL
- username VARCHAR(100) NOT NULL
- password_hash VARCHAR(255) NOT NULL
- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
- updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

Indexes: idx_users_email

Table: chats
- id UUID PRIMARY KEY DEFAULT uuid_generate_v4()
- user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
- title VARCHAR(255) NOT NULL
- description TEXT
- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
- updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

Indexes: idx_chats_user_id, idx_chats_updated_at

Table: messages
- id UUID PRIMARY KEY DEFAULT uuid_generate_v4()
- chat_id UUID NOT NULL REFERENCES chats(id) ON DELETE CASCADE
- content TEXT NOT NULL
- role VARCHAR(20) NOT NULL CHECK (role IN ('user','assistant','system'))
- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
- updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

Indexes: idx_messages_chat_id, idx_messages_created_at

## Migrations & Extensions
- The initial migration enables the PostgreSQL `uuid-ossp` extension and uses `uuid_generate_v4()` for server-side generated UUIDs.
- Migration files are in `migrations/` and are applied automatically at server startup by `internal/database.RunMigrations`.

## Environment variables (from `.env.example`)
- DATABASE_URL: `postgres://user:pass@host:port/dbname?sslmode=disable`
- JWT_SECRET: secret used to sign JWTs
- PORT: server port (default 8080)
- GIN_MODE: gin mode (debug/release)
- OPENAI_API_KEY / ANTHROPIC_API_KEY: placeholders for future LLM provider keys

## Authentication & Tokens
- JWT tokens generated by `UserService.GenerateTokens` (access: 15m, refresh: 7d in current code)
- `AuthMiddleware` validates tokens and populates `userID` and `email` into Gin context
- Secret used is `JWT_SECRET` from env (code currently uses TODO placeholders; replace with `config.JWTSecret` for production)

## Error handling recommendations (future)
- Introduce typed domain errors (e.g., ErrNotFound, ErrUnauthorized) to enable programmatic translation to HTTP codes.
- Centralize error to HTTP mapping in a single helper to avoid duplication.

## Quick `curl` examples

Register:

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","username":"user","password":"password123"}'
```

Create chat (replace TOKEN):

```bash
curl -X POST http://localhost:8080/api/v1/chats \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"My chat","description":"Notes"}'
```

Send message:

```bash
curl -X POST http://localhost:8080/api/v1/chats/<chat-id>/messages \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"content":"Hello","role":"user"}'
```

## Developer checklist to implement/validate
- [ ] Ensure `DATABASE_URL` is reachable and Postgres server is running
- [ ] Verify migrations run and `uuid-ossp` is enabled
- [ ] Create test users, chats, messages via API
- [ ] Replace JWT placeholder secret usage with `config.JWTSecret`
- [ ] Add integration tests for key endpoints


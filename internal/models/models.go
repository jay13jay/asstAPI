package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Chat represents a chat conversation
type Chat struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Message represents a message in a chat
type Message struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	ChatID    uuid.UUID   `json:"chat_id" db:"chat_id"`
	Content   string      `json:"content" db:"content"`
	Role      MessageRole `json:"role" db:"role"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}

// MessageRole represents the role of a message sender
type MessageRole string

const (
	MessageRoleUser      MessageRole = "user"
	MessageRoleAssistant MessageRole = "assistant"
	MessageRoleSystem    MessageRole = "system"
)

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// CreateChatRequest represents the request to create a new chat
type CreateChatRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=255"`
	Description string `json:"description,omitempty"`
}

// UpdateChatRequest represents the request to update a chat
type UpdateChatRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateMessageRequest represents the request to create a new message
type CreateMessageRequest struct {
	Content string      `json:"content" binding:"required"`
	Role    MessageRole `json:"role" binding:"required,oneof=user assistant system"`
}

// UpdateMessageRequest represents the request to update a message
type UpdateMessageRequest struct {
	Content string `json:"content" binding:"required"`
}

// AuthResponse represents the authentication response
type AuthResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Claims represents JWT claims
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

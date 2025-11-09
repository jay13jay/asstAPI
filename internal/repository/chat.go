package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jay13jay/asstBackend/internal/models"
)

// ChatRepository handles chat data operations
type ChatRepository struct {
	db *sql.DB
}

// NewChatRepository creates a new ChatRepository
func NewChatRepository(db *sql.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

// Create creates a new chat
func (r *ChatRepository) Create(chat *models.Chat) (*models.Chat, error) {
	query := `
		INSERT INTO chats (id, user_id, title, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, user_id, title, description, created_at, updated_at`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		chat.ID,
		chat.UserID,
		chat.Title,
		chat.Description,
		now,
		now,
	).Scan(
		&chat.ID,
		&chat.UserID,
		&chat.Title,
		&chat.Description,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return chat, nil
}

// GetByID retrieves a chat by ID
func (r *ChatRepository) GetByID(id uuid.UUID) (*models.Chat, error) {
	query := `
		SELECT id, user_id, title, description, created_at, updated_at
		FROM chats WHERE id = $1`

	chat := &models.Chat{}
	err := r.db.QueryRow(query, id).Scan(
		&chat.ID,
		&chat.UserID,
		&chat.Title,
		&chat.Description,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get chat by ID: %w", err)
	}

	return chat, nil
}

// GetByUserID retrieves all chats for a user
func (r *ChatRepository) GetByUserID(userID uuid.UUID) ([]*models.Chat, error) {
	query := `
		SELECT id, user_id, title, description, created_at, updated_at
		FROM chats WHERE user_id = $1 ORDER BY updated_at DESC`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chats by user ID: %w", err)
	}
	defer rows.Close()

	var chats []*models.Chat
	for rows.Next() {
		chat := &models.Chat{}
		err := rows.Scan(
			&chat.ID,
			&chat.UserID,
			&chat.Title,
			&chat.Description,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan chat: %w", err)
		}
		chats = append(chats, chat)
	}

	return chats, nil
}

// Update updates a chat
func (r *ChatRepository) Update(chat *models.Chat) (*models.Chat, error) {
	query := `
		UPDATE chats 
		SET title = $2, description = $3, updated_at = $4
		WHERE id = $1
		RETURNING id, user_id, title, description, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		chat.ID,
		chat.Title,
		chat.Description,
		time.Now(),
	).Scan(
		&chat.ID,
		&chat.UserID,
		&chat.Title,
		&chat.Description,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update chat: %w", err)
	}

	return chat, nil
}

// Delete deletes a chat
func (r *ChatRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM chats WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}

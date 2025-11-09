package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jay13jay/asstBackend/internal/models"
)

// MessageRepository handles message data operations
type MessageRepository struct {
	db *sql.DB
}

// NewMessageRepository creates a new MessageRepository
func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// Create creates a new message
func (r *MessageRepository) Create(message *models.Message) (*models.Message, error) {
	query := `
		INSERT INTO messages (id, chat_id, content, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, chat_id, content, role, created_at, updated_at`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		message.ID,
		message.ChatID,
		message.Content,
		message.Role,
		now,
		now,
	).Scan(
		&message.ID,
		&message.ChatID,
		&message.Content,
		&message.Role,
		&message.CreatedAt,
		&message.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}

	return message, nil
}

// GetByID retrieves a message by ID
func (r *MessageRepository) GetByID(id uuid.UUID) (*models.Message, error) {
	query := `
		SELECT id, chat_id, content, role, created_at, updated_at
		FROM messages WHERE id = $1`

	message := &models.Message{}
	err := r.db.QueryRow(query, id).Scan(
		&message.ID,
		&message.ChatID,
		&message.Content,
		&message.Role,
		&message.CreatedAt,
		&message.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get message by ID: %w", err)
	}

	return message, nil
}

// GetByChatID retrieves all messages for a chat
func (r *MessageRepository) GetByChatID(chatID uuid.UUID) ([]*models.Message, error) {
	query := `
		SELECT id, chat_id, content, role, created_at, updated_at
		FROM messages WHERE chat_id = $1 ORDER BY created_at ASC`

	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages by chat ID: %w", err)
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		message := &models.Message{}
		err := rows.Scan(
			&message.ID,
			&message.ChatID,
			&message.Content,
			&message.Role,
			&message.CreatedAt,
			&message.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, message)
	}

	return messages, nil
}

// Update updates a message
func (r *MessageRepository) Update(message *models.Message) (*models.Message, error) {
	query := `
		UPDATE messages 
		SET content = $2, updated_at = $3
		WHERE id = $1
		RETURNING id, chat_id, content, role, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		message.ID,
		message.Content,
		time.Now(),
	).Scan(
		&message.ID,
		&message.ChatID,
		&message.Content,
		&message.Role,
		&message.CreatedAt,
		&message.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update message: %w", err)
	}

	return message, nil
}

// Delete deletes a message
func (r *MessageRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM messages WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	return nil
}

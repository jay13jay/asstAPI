package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jay13jay/asstBackend/internal/models"
	"github.com/jay13jay/asstBackend/internal/repository"
)

// MessageService provides message-related business logic
type MessageService struct {
	messageRepo *repository.MessageRepository
}

// NewMessageService creates a new MessageService
func NewMessageService(messageRepo *repository.MessageRepository) *MessageService {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

// CreateMessage creates a new message
func (s *MessageService) CreateMessage(chatID uuid.UUID, req *models.CreateMessageRequest) (*models.Message, error) {
	message := &models.Message{
		ID:      uuid.New(),
		ChatID:  chatID,
		Content: req.Content,
		Role:    req.Role,
	}

	createdMessage, err := s.messageRepo.Create(message)
	if err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}

	return createdMessage, nil
}

// GetMessage retrieves a message by ID
func (s *MessageService) GetMessage(messageID uuid.UUID) (*models.Message, error) {
	message, err := s.messageRepo.GetByID(messageID)
	if err != nil {
		return nil, fmt.Errorf("message not found")
	}

	return message, nil
}

// ListMessages retrieves all messages for a chat
func (s *MessageService) ListMessages(chatID uuid.UUID) ([]*models.Message, error) {
	messages, err := s.messageRepo.GetByChatID(chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to list messages: %w", err)
	}

	return messages, nil
}

// UpdateMessage updates a message
func (s *MessageService) UpdateMessage(messageID uuid.UUID, req *models.UpdateMessageRequest) (*models.Message, error) {
	message, err := s.GetMessage(messageID)
	if err != nil {
		return nil, err
	}

	message.Content = req.Content

	updatedMessage, err := s.messageRepo.Update(message)
	if err != nil {
		return nil, fmt.Errorf("failed to update message: %w", err)
	}

	return updatedMessage, nil
}

// DeleteMessage deletes a message
func (s *MessageService) DeleteMessage(messageID uuid.UUID) error {
	err := s.messageRepo.Delete(messageID)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}

	return nil
}

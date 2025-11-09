package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jay13jay/asstBackend/internal/models"
	"github.com/jay13jay/asstBackend/internal/repository"
)

// ChatService provides chat-related business logic
type ChatService struct {
	chatRepo *repository.ChatRepository
}

// NewChatService creates a new ChatService
func NewChatService(chatRepo *repository.ChatRepository) *ChatService {
	return &ChatService{
		chatRepo: chatRepo,
	}
}

// CreateChat creates a new chat
func (s *ChatService) CreateChat(userID uuid.UUID, req *models.CreateChatRequest) (*models.Chat, error) {
	chat := &models.Chat{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
	}

	createdChat, err := s.chatRepo.Create(chat)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return createdChat, nil
}

// GetChat retrieves a chat by ID
func (s *ChatService) GetChat(chatID, userID uuid.UUID) (*models.Chat, error) {
	chat, err := s.chatRepo.GetByID(chatID)
	if err != nil {
		return nil, fmt.Errorf("chat not found")
	}

	// Ensure user owns this chat
	if chat.UserID != userID {
		return nil, fmt.Errorf("unauthorized access to chat")
	}

	return chat, nil
}

// ListChats retrieves all chats for a user
func (s *ChatService) ListChats(userID uuid.UUID) ([]*models.Chat, error) {
	chats, err := s.chatRepo.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list chats: %w", err)
	}

	return chats, nil
}

// UpdateChat updates a chat
func (s *ChatService) UpdateChat(chatID, userID uuid.UUID, req *models.UpdateChatRequest) (*models.Chat, error) {
	// First verify ownership
	chat, err := s.GetChat(chatID, userID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Title != "" {
		chat.Title = req.Title
	}
	if req.Description != "" {
		chat.Description = req.Description
	}

	updatedChat, err := s.chatRepo.Update(chat)
	if err != nil {
		return nil, fmt.Errorf("failed to update chat: %w", err)
	}

	return updatedChat, nil
}

// DeleteChat deletes a chat
func (s *ChatService) DeleteChat(chatID, userID uuid.UUID) error {
	// First verify ownership
	_, err := s.GetChat(chatID, userID)
	if err != nil {
		return err
	}

	err = s.chatRepo.Delete(chatID)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}

package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jay13jay/asstBackend/internal/models"
)

// UserRepository handles user data operations
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (id, email, username, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, email, username, created_at, updated_at`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		user.ID,
		user.Email,
		user.Username,
		user.Password,
		now,
		now,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Clear password from returned user
	user.Password = ""
	return user, nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, email, username, password_hash, created_at, updated_at
		FROM users WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, email, username, password_hash, created_at, updated_at
		FROM users WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

// Update updates a user
func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	query := `
		UPDATE users 
		SET username = $2, updated_at = $3
		WHERE id = $1
		RETURNING id, email, username, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		user.ID,
		user.Username,
		time.Now(),
	).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

// Delete deletes a user
func (r *UserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

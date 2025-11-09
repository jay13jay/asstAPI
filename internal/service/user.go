package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jay13jay/asstBackend/internal/models"
	"github.com/jay13jay/asstBackend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService provides user-related business logic
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}

	createdUser, err := s.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return createdUser, nil
}

// AuthenticateUser verifies user credentials
func (s *UserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

// GenerateTokens generates access and refresh tokens for a user
func (s *UserService) GenerateTokens(user *models.User) (string, string, error) {
	// Access token (expires in 15 minutes)
	accessClaims := &models.Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte("your-secret-key")) // TODO: Use config
	if err != nil {
		return "", "", err
	}

	// Refresh token (expires in 7 days)
	refreshClaims := &models.Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("your-secret-key")) // TODO: Use config
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

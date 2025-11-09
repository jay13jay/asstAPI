package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jay13jay/asstBackend/internal/config"
	"github.com/jay13jay/asstBackend/internal/handlers"
	"github.com/jay13jay/asstBackend/internal/middleware"
	"github.com/jay13jay/asstBackend/internal/repository"
	"github.com/jay13jay/asstBackend/internal/service"
)

// Server represents the API server
type Server struct {
	router *gin.Engine
	db     *sql.DB
	config *config.Config
}

// NewServer creates a new API server instance
func NewServer(db *sql.DB, cfg *config.Config) *Server {
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	server := &Server{
		router: router,
		db:     db,
		config: cfg,
	}

	server.setupRoutes()
	return server
}

// setupRoutes configures all API routes
func (s *Server) setupRoutes() {
	// Initialize repositories
	userRepo := repository.NewUserRepository(s.db)
	chatRepo := repository.NewChatRepository(s.db)
	messageRepo := repository.NewMessageRepository(s.db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	chatService := service.NewChatService(chatRepo)
	messageService := service.NewMessageService(messageRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	chatHandler := handlers.NewChatHandler(chatService)
	messageHandler := handlers.NewMessageHandler(messageService)

	// Health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	{
		// Authentication routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh", userHandler.RefreshToken)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware(s.config.JWTSecret))
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/me", userHandler.GetProfile)
				users.PUT("/me", userHandler.UpdateProfile)
			}

			// Chat routes
			chats := protected.Group("/chats")
			{
				chats.POST("/", chatHandler.CreateChat)
				chats.GET("/", chatHandler.ListChats)
				chats.GET("/:id", chatHandler.GetChat)
				chats.PUT("/:id", chatHandler.UpdateChat)
				chats.DELETE("/:id", chatHandler.DeleteChat)

				// Message routes within chats
				chats.POST("/:id/messages", messageHandler.CreateMessage)
				chats.GET("/:id/messages", messageHandler.ListMessages)
				chats.PUT("/:id/messages/:messageId", messageHandler.UpdateMessage)
				chats.DELETE("/:id/messages/:messageId", messageHandler.DeleteMessage)
			}
		}
	}
}

// Run starts the server
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

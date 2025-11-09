package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// EnsureDatabaseExists checks if the database exists and creates it if not
func EnsureDatabaseExists(databaseURL string) error {
	// Parse the database URL to extract database name
	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return fmt.Errorf("failed to parse database URL: %w", err)
	}

	// Extract database name from path (removing leading slash)
	dbName := strings.TrimPrefix(parsedURL.Path, "/")
	if dbName == "" {
		return fmt.Errorf("database name not found in URL")
	}

	// Create connection URL to 'postgres' system database
	systemURL := *parsedURL
	systemURL.Path = "/postgres"
	systemDatabaseURL := systemURL.String()

	// Connect to system database
	systemDB, err := sql.Open("postgres", systemDatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to connect to system database: %w", err)
	}
	defer systemDB.Close()

	// Check if target database exists
	var exists bool
	query := "SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = $1)"
	err = systemDB.QueryRow(query, dbName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	// Create database if it doesn't exist
	if !exists {
		createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
		_, err = systemDB.Exec(createQuery)
		if err != nil {
			return fmt.Errorf("failed to create database %s: %w", dbName, err)
		}
		fmt.Printf("Database '%s' created successfully\n", dbName)
	} else {
		fmt.Printf("Database '%s' already exists\n", dbName)
	}

	return nil
}

// NewConnection creates a new database connection
func NewConnection(databaseURL string) (*sql.DB, error) {
	// Ensure database exists before connecting
	if err := EnsureDatabaseExists(databaseURL); err != nil {
		return nil, fmt.Errorf("failed to ensure database exists: %w", err)
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// RunMigrations runs database migrations
func RunMigrations(databaseURL string) error {
	// Ensure database exists before running migrations
	if err := EnsureDatabaseExists(databaseURL); err != nil {
		return fmt.Errorf("failed to ensure database exists: %w", err)
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("failed to open database for migrations: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

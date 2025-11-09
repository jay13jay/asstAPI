package database

import (
	"net/url"
	"strings"
	"testing"
)

func TestParseDatabaseURL(t *testing.T) {
	testCases := []struct {
		name        string
		databaseURL string
		expectedDB  string
		shouldError bool
	}{
		{
			name:        "Valid PostgreSQL URL",
			databaseURL: "postgres://user:pass@localhost:5432/testdb?sslmode=disable",
			expectedDB:  "testdb",
			shouldError: false,
		},
		{
			name:        "PostgreSQL URL without credentials",
			databaseURL: "postgres://localhost:5432/asstbackend",
			expectedDB:  "asstbackend",
			shouldError: false,
		},
		{
			name:        "URL with special characters in database name",
			databaseURL: "postgres://localhost:5432/test_db_123",
			expectedDB:  "test_db_123",
			shouldError: false,
		},
		{
			name:        "Invalid URL",
			databaseURL: "not-a-valid-url",
			expectedDB:  "",
			shouldError: true,
		},
		{
			name:        "URL without database name",
			databaseURL: "postgres://localhost:5432/",
			expectedDB:  "",
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// This mimics the parsing logic in EnsureDatabaseExists
			parsedURL, err := url.Parse(tc.databaseURL)
			if err != nil {
				if !tc.shouldError {
					t.Errorf("Unexpected error parsing URL: %v", err)
				}
				return
			}

			// url.Parse does not always return an error for invalid-looking URLs
			// Treat URLs missing a scheme or host as invalid for our purposes.
			if parsedURL.Scheme == "" || parsedURL.Host == "" {
				if !tc.shouldError {
					t.Errorf("Unexpected invalid URL (missing scheme or host): %v", tc.databaseURL)
				}
				return
			}

			dbName := strings.TrimPrefix(parsedURL.Path, "/")
			if dbName == "" && !tc.shouldError {
				t.Error("Expected database name but got empty string")
				return
			}

			if tc.shouldError && dbName != "" {
				t.Error("Expected error but got valid database name")
				return
			}

			if dbName != tc.expectedDB {
				t.Errorf("Expected database name '%s', got '%s'", tc.expectedDB, dbName)
			}
		})
	}
}

func TestSystemDatabaseURL(t *testing.T) {
	databaseURL := "postgres://user:pass@localhost:5432/testdb?sslmode=disable"

	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		t.Fatalf("Failed to parse URL: %v", err)
	}

	// Create system database URL (connects to 'postgres' system database)
	systemURL := *parsedURL
	systemURL.Path = "/postgres"

	expectedSystemURL := "postgres://user:pass@localhost:5432/postgres?sslmode=disable"
	actualSystemURL := systemURL.String()

	if actualSystemURL != expectedSystemURL {
		t.Errorf("Expected system URL '%s', got '%s'", expectedSystemURL, actualSystemURL)
	}
}

package storage

import (
	"errors"
	"url-shortener/models"
)

var (
	ErrNotFound      = errors.New("URL not found")
	ErrAlreadyExists = errors.New("short code already exists")
)

// Storage defines the interface for URL storage operations
type Storage interface {
	// Save stores a new URL mapping
	Save(url *models.URL) error

	// Get retrieves a URL by short code
	Get(shortCode string) (*models.URL, error)

	// Update updates an existing URL
	Update(url *models.URL) error

	// Delete removes a URL by short code
	Delete(shortCode string) error

	// List returns all URLs (for admin purposes)
	List(limit, offset int) ([]*models.URL, error)

	// Close closes any database connections
	Close() error
}

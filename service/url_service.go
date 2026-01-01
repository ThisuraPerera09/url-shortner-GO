package service

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"
	"url-shortener/models"
	"url-shortener/storage"
)

// URLService handles business logic for URL shortening
type URLService struct {
	storage      storage.Storage
	shortCodeLen int
}

func NewURLService(storage storage.Storage, shortCodeLen int) *URLService {
	return &URLService{
		storage:      storage,
		shortCodeLen: shortCodeLen,
	}
}

// ShortenURL creates a short code for the given URL
func (s *URLService) ShortenURL(originalURL, customCode string) (*models.URL, error) {
	var shortCode string
	var err error

	if customCode != "" {
		// Use custom code if provided
		shortCode = customCode
	} else {
		// Generate random short code
		shortCode, err = s.generateShortCode()
		if err != nil {
			return nil, err
		}
	}

	url := &models.URL{
		ShortCode:   shortCode,
		OriginalURL: originalURL,
	}

	// Try to save, if collision occurs, try again (only for generated codes)
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		err = s.storage.Save(url)
		if err == nil {
			return url, nil
		}

		if err == storage.ErrAlreadyExists && customCode == "" {
			// Generate new code and retry
			shortCode, err = s.generateShortCode()
			if err != nil {
				return nil, err
			}
			url.ShortCode = shortCode
			continue
		}

		return nil, err
	}

	return nil, err
}

// GetURL retrieves the original URL and increments click count
func (s *URLService) GetURL(shortCode string) (*models.URL, error) {
	url, err := s.storage.Get(shortCode)
	if err != nil {
		return nil, err
	}

	// Increment click count
	url.Clicks++
	now := time.Now()
	url.LastAccessed = &now

	// Update in background (we don't want to slow down the redirect)
	go s.storage.Update(url)

	return url, nil
}

// GetStats retrieves URL statistics without incrementing click count
func (s *URLService) GetStats(shortCode string) (*models.URL, error) {
	return s.storage.Get(shortCode)
}

// DeleteURL removes a shortened URL
func (s *URLService) DeleteURL(shortCode string) error {
	return s.storage.Delete(shortCode)
}

// ListURLs retrieves all URLs with pagination
func (s *URLService) ListURLs(limit, offset int) ([]*models.URL, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.storage.List(limit, offset)
}

// generateShortCode creates a random URL-safe short code
func (s *URLService) generateShortCode() (string, error) {
	// Generate random bytes
	b := make([]byte, s.shortCodeLen)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Encode to base64 and make URL-safe
	code := base64.URLEncoding.EncodeToString(b)
	code = strings.TrimRight(code, "=")

	// Take only the desired length
	if len(code) > s.shortCodeLen {
		code = code[:s.shortCodeLen]
	}

	return code, nil
}

// Close closes the storage connection
func (s *URLService) Close() error {
	return s.storage.Close()
}

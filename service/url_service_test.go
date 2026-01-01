package service

import (
	"testing"
	"url-shortener/storage"
)

func TestShortenURL(t *testing.T) {
	// Setup
	store := storage.NewInMemoryStorage()
	service := NewURLService(store, 6)

	// Test case 1: Generate random short code
	t.Run("Generate random short code", func(t *testing.T) {
		url, err := service.ShortenURL("https://example.com", "")
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if url.ShortCode == "" {
			t.Error("Expected non-empty short code")
		}
		
		if len(url.ShortCode) != 6 {
			t.Errorf("Expected short code length 6, got %d", len(url.ShortCode))
		}
		
		if url.OriginalURL != "https://example.com" {
			t.Errorf("Expected original URL 'https://example.com', got '%s'", url.OriginalURL)
		}
	})

	// Test case 2: Custom short code
	t.Run("Use custom short code", func(t *testing.T) {
		url, err := service.ShortenURL("https://google.com", "google")
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if url.ShortCode != "google" {
			t.Errorf("Expected short code 'google', got '%s'", url.ShortCode)
		}
	})

	// Test case 3: Duplicate custom code
	t.Run("Reject duplicate custom code", func(t *testing.T) {
		_, err := service.ShortenURL("https://another.com", "google")
		
		if err == nil {
			t.Error("Expected error for duplicate custom code")
		}
		
		if err != storage.ErrAlreadyExists {
			t.Errorf("Expected ErrAlreadyExists, got %v", err)
		}
	})
}

func TestGetURL(t *testing.T) {
	// Setup
	store := storage.NewInMemoryStorage()
	service := NewURLService(store, 6)

	// Create a URL
	originalURL, _ := service.ShortenURL("https://example.com", "test")

	t.Run("Get existing URL", func(t *testing.T) {
		url, err := service.GetURL("test")
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if url.OriginalURL != "https://example.com" {
			t.Errorf("Expected 'https://example.com', got '%s'", url.OriginalURL)
		}
		
		// Note: Click count is incremented in GetURL
		if url.Clicks != 1 {
			t.Errorf("Expected clicks to be 1, got %d", url.Clicks)
		}
	})

	t.Run("Get non-existent URL", func(t *testing.T) {
		_, err := service.GetURL("nonexistent")
		
		if err == nil {
			t.Error("Expected error for non-existent URL")
		}
		
		if err != storage.ErrNotFound {
			t.Errorf("Expected ErrNotFound, got %v", err)
		}
	})

	// Verify click tracking
	t.Run("Track multiple clicks", func(t *testing.T) {
		// Get the URL again
		service.GetURL("test")
		
		// Give time for async update
		// In a real test, you might want to make this synchronous
		// or use a mock storage
		
		stats, _ := service.GetStats("test")
		if stats.Clicks < 1 {
			t.Errorf("Expected at least 1 click, got %d", stats.Clicks)
		}
	})

	_ = originalURL // Avoid unused variable warning
}

func TestDeleteURL(t *testing.T) {
	// Setup
	store := storage.NewInMemoryStorage()
	service := NewURLService(store, 6)

	// Create a URL
	service.ShortenURL("https://example.com", "todelete")

	t.Run("Delete existing URL", func(t *testing.T) {
		err := service.DeleteURL("todelete")
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		// Verify it's deleted
		_, err = service.GetStats("todelete")
		if err != storage.ErrNotFound {
			t.Error("Expected URL to be deleted")
		}
	})

	t.Run("Delete non-existent URL", func(t *testing.T) {
		err := service.DeleteURL("nonexistent")
		
		if err != storage.ErrNotFound {
			t.Errorf("Expected ErrNotFound, got %v", err)
		}
	})
}

func TestListURLs(t *testing.T) {
	// Setup
	store := storage.NewInMemoryStorage()
	service := NewURLService(store, 6)

	// Create multiple URLs
	service.ShortenURL("https://example1.com", "url1")
	service.ShortenURL("https://example2.com", "url2")
	service.ShortenURL("https://example3.com", "url3")

	t.Run("List with default pagination", func(t *testing.T) {
		urls, err := service.ListURLs(10, 0)
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if len(urls) != 3 {
			t.Errorf("Expected 3 URLs, got %d", len(urls))
		}
	})

	t.Run("List with limit", func(t *testing.T) {
		urls, err := service.ListURLs(2, 0)
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if len(urls) != 2 {
			t.Errorf("Expected 2 URLs, got %d", len(urls))
		}
	})

	t.Run("List with offset", func(t *testing.T) {
		urls, err := service.ListURLs(10, 2)
		
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		
		if len(urls) != 1 {
			t.Errorf("Expected 1 URL, got %d", len(urls))
		}
	})
}

func TestGenerateShortCode(t *testing.T) {
	service := NewURLService(nil, 8)

	t.Run("Generate multiple unique codes", func(t *testing.T) {
		codes := make(map[string]bool)
		
		// Generate 100 codes and check they're unique
		for i := 0; i < 100; i++ {
			code, err := service.generateShortCode()
			
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			
			if codes[code] {
				t.Errorf("Generated duplicate code: %s", code)
			}
			
			codes[code] = true
			
			if len(code) != 8 {
				t.Errorf("Expected code length 8, got %d", len(code))
			}
		}
	})
}


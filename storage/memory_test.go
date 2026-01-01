package storage

import (
	"testing"
	"url-shortener/models"
)

func TestInMemoryStorage(t *testing.T) {
	store := NewInMemoryStorage()

	t.Run("Save and Get URL", func(t *testing.T) {
		url := &models.URL{
			ShortCode:   "test123",
			OriginalURL: "https://example.com",
		}

		// Save
		err := store.Save(url)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Get
		retrieved, err := store.Get("test123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if retrieved.ShortCode != "test123" {
			t.Errorf("Expected short code 'test123', got '%s'", retrieved.ShortCode)
		}

		if retrieved.OriginalURL != "https://example.com" {
			t.Errorf("Expected URL 'https://example.com', got '%s'", retrieved.OriginalURL)
		}
	})

	t.Run("Prevent duplicate short codes", func(t *testing.T) {
		url := &models.URL{
			ShortCode:   "duplicate",
			OriginalURL: "https://example1.com",
		}
		store.Save(url)

		// Try to save duplicate
		duplicate := &models.URL{
			ShortCode:   "duplicate",
			OriginalURL: "https://example2.com",
		}
		err := store.Save(duplicate)

		if err != ErrAlreadyExists {
			t.Errorf("Expected ErrAlreadyExists, got %v", err)
		}
	})

	t.Run("Update URL", func(t *testing.T) {
		url := &models.URL{
			ShortCode:   "update",
			OriginalURL: "https://example.com",
			Clicks:      0,
		}
		store.Save(url)

		// Update
		url.Clicks = 5
		err := store.Update(url)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Verify
		retrieved, _ := store.Get("update")
		if retrieved.Clicks != 5 {
			t.Errorf("Expected 5 clicks, got %d", retrieved.Clicks)
		}
	})

	t.Run("Delete URL", func(t *testing.T) {
		url := &models.URL{
			ShortCode:   "delete",
			OriginalURL: "https://example.com",
		}
		store.Save(url)

		// Delete
		err := store.Delete("delete")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Verify deleted
		_, err = store.Get("delete")
		if err != ErrNotFound {
			t.Error("Expected ErrNotFound after deletion")
		}
	})

	t.Run("List URLs with pagination", func(t *testing.T) {
		// Create fresh store
		store := NewInMemoryStorage()
		
		// Add multiple URLs
		for i := 0; i < 5; i++ {
			url := &models.URL{
				ShortCode:   string(rune('a' + i)),
				OriginalURL: "https://example.com",
			}
			store.Save(url)
		}

		// List all
		urls, err := store.List(10, 0)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if len(urls) != 5 {
			t.Errorf("Expected 5 URLs, got %d", len(urls))
		}

		// List with limit
		urls, _ = store.List(2, 0)
		if len(urls) != 2 {
			t.Errorf("Expected 2 URLs, got %d", len(urls))
		}

		// List with offset
		urls, _ = store.List(10, 3)
		if len(urls) != 2 {
			t.Errorf("Expected 2 URLs, got %d", len(urls))
		}
	})
}


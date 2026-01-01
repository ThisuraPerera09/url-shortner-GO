package models

import "time"

// URL represents a shortened URL entry
type URL struct {
	ID           int64      `json:"id"`
	ShortCode    string     `json:"short_code"`
	OriginalURL  string     `json:"original_url"`
	Clicks       int64      `json:"clicks"`
	CreatedAt    time.Time  `json:"created_at"`
	LastAccessed *time.Time `json:"last_accessed,omitempty"`
}

// ShortenRequest represents the request to shorten a URL
type ShortenRequest struct {
	URL        string `json:"url" binding:"required,url"`
	CustomCode string `json:"custom_code,omitempty"`
}

// ShortenResponse represents the response after shortening a URL
type ShortenResponse struct {
	ShortCode   string `json:"short_code"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// StatsResponse represents URL statistics
type StatsResponse struct {
	ShortCode    string     `json:"short_code"`
	OriginalURL  string     `json:"original_url"`
	Clicks       int64      `json:"clicks"`
	CreatedAt    time.Time  `json:"created_at"`
	LastAccessed *time.Time `json:"last_accessed,omitempty"`
}

package storage

import (
	"sync"
	"time"
	"url-shortener/models"
)

// InMemoryStorage implements Storage interface using a map
type InMemoryStorage struct {
	urls      map[string]*models.URL
	mutex     sync.RWMutex
	idCounter int64
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		urls:      make(map[string]*models.URL),
		idCounter: 0,
	}
}

func (s *InMemoryStorage) Save(url *models.URL) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.urls[url.ShortCode]; exists {
		return ErrAlreadyExists
	}

	s.idCounter++
	url.ID = s.idCounter
	url.CreatedAt = time.Now()
	url.Clicks = 0

	s.urls[url.ShortCode] = url
	return nil
}

func (s *InMemoryStorage) Get(shortCode string) (*models.URL, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	url, exists := s.urls[shortCode]
	if !exists {
		return nil, ErrNotFound
	}

	return url, nil
}

func (s *InMemoryStorage) Update(url *models.URL) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.urls[url.ShortCode]; !exists {
		return ErrNotFound
	}

	s.urls[url.ShortCode] = url
	return nil
}

func (s *InMemoryStorage) Delete(shortCode string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.urls[shortCode]; !exists {
		return ErrNotFound
	}

	delete(s.urls, shortCode)
	return nil
}

func (s *InMemoryStorage) List(limit, offset int) ([]*models.URL, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	urls := make([]*models.URL, 0, len(s.urls))
	for _, url := range s.urls {
		urls = append(urls, url)
	}

	// Apply offset and limit
	start := offset
	if start > len(urls) {
		return []*models.URL{}, nil
	}

	end := start + limit
	if end > len(urls) {
		end = len(urls)
	}

	return urls[start:end], nil
}

func (s *InMemoryStorage) Close() error {
	return nil
}

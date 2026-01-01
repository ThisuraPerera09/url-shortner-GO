package storage

import (
	"database/sql"
	"time"
	"url-shortener/models"

	_ "github.com/lib/pq"
)

// PostgresStorage implements Storage interface using PostgreSQL
type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connectionString string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	storage := &PostgresStorage{db: db}
	if err := storage.createTable(); err != nil {
		db.Close()
		return nil, err
	}

	return storage, nil
}

func (s *PostgresStorage) createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		short_code VARCHAR(255) UNIQUE NOT NULL,
		original_url TEXT NOT NULL,
		clicks BIGINT DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		last_accessed TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_short_code ON urls(short_code);
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) Save(url *models.URL) error {
	query := `INSERT INTO urls (short_code, original_url, clicks, created_at) 
	          VALUES ($1, $2, $3, $4) RETURNING id`

	err := s.db.QueryRow(query, url.ShortCode, url.OriginalURL, 0, time.Now()).Scan(&url.ID)
	if err != nil {
		// Check if it's a unique constraint error
		if err.Error() == "pq: duplicate key value violates unique constraint \"urls_short_code_key\"" {
			return ErrAlreadyExists
		}
		return err
	}

	return nil
}

func (s *PostgresStorage) Get(shortCode string) (*models.URL, error) {
	query := `SELECT id, short_code, original_url, clicks, created_at, last_accessed 
	          FROM urls WHERE short_code = $1`

	url := &models.URL{}
	var lastAccessed sql.NullTime

	err := s.db.QueryRow(query, shortCode).Scan(
		&url.ID,
		&url.ShortCode,
		&url.OriginalURL,
		&url.Clicks,
		&url.CreatedAt,
		&lastAccessed,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	if lastAccessed.Valid {
		url.LastAccessed = &lastAccessed.Time
	}

	return url, nil
}

func (s *PostgresStorage) Update(url *models.URL) error {
	query := `UPDATE urls SET original_url = $1, clicks = $2, last_accessed = $3 
	          WHERE short_code = $4`

	result, err := s.db.Exec(query, url.OriginalURL, url.Clicks, url.LastAccessed, url.ShortCode)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostgresStorage) Delete(shortCode string) error {
	query := `DELETE FROM urls WHERE short_code = $1`

	result, err := s.db.Exec(query, shortCode)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostgresStorage) List(limit, offset int) ([]*models.URL, error) {
	query := `SELECT id, short_code, original_url, clicks, created_at, last_accessed 
	          FROM urls ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*models.URL
	for rows.Next() {
		url := &models.URL{}
		var lastAccessed sql.NullTime

		err := rows.Scan(
			&url.ID,
			&url.ShortCode,
			&url.OriginalURL,
			&url.Clicks,
			&url.CreatedAt,
			&lastAccessed,
		)
		if err != nil {
			return nil, err
		}

		if lastAccessed.Valid {
			url.LastAccessed = &lastAccessed.Time
		}

		urls = append(urls, url)
	}

	return urls, rows.Err()
}

func (s *PostgresStorage) Close() error {
	return s.db.Close()
}


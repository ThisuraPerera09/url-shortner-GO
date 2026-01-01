package storage

import (
	"database/sql"
	"time"
	"url-shortener/models"

	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStorage implements Storage interface using SQLite
type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	storage := &SQLiteStorage{db: db}
	if err := storage.createTable(); err != nil {
		db.Close()
		return nil, err
	}

	return storage, nil
}

func (s *SQLiteStorage) createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_code TEXT UNIQUE NOT NULL,
		original_url TEXT NOT NULL,
		clicks INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		last_accessed DATETIME
	);
	CREATE INDEX IF NOT EXISTS idx_short_code ON urls(short_code);
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *SQLiteStorage) Save(url *models.URL) error {
	query := `INSERT INTO urls (short_code, original_url, clicks, created_at) VALUES (?, ?, ?, ?)`

	result, err := s.db.Exec(query, url.ShortCode, url.OriginalURL, 0, time.Now())
	if err != nil {
		// Check if it's a unique constraint error
		if err.Error() == "UNIQUE constraint failed: urls.short_code" {
			return ErrAlreadyExists
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	url.ID = id
	return nil
}

func (s *SQLiteStorage) Get(shortCode string) (*models.URL, error) {
	query := `SELECT id, short_code, original_url, clicks, created_at, last_accessed FROM urls WHERE short_code = ?`

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

func (s *SQLiteStorage) Update(url *models.URL) error {
	query := `UPDATE urls SET original_url = ?, clicks = ?, last_accessed = ? WHERE short_code = ?`

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

func (s *SQLiteStorage) Delete(shortCode string) error {
	query := `DELETE FROM urls WHERE short_code = ?`

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

func (s *SQLiteStorage) List(limit, offset int) ([]*models.URL, error) {
	query := `SELECT id, short_code, original_url, clicks, created_at, last_accessed 
	          FROM urls ORDER BY created_at DESC LIMIT ? OFFSET ?`

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

func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}

package repository

import (
	"awura-shortener/internal/model"
	"database/sql"
)

type URLRepository struct {
	DB *sql.DB
}

func (r *URLRepository) CreateURL(url *model.URL) error {
	query := `INSERT INTO urls (short_url, original_url, created_at) VALUES ($1, $2, NOW())`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(url.ShortURL, url.OriginalURL)
	return err
}

func (r *URLRepository) GetURL(shortURL string) (*model.URL, error) {
	query := `SELECT id, short_url, original_url, created_at FROM urls WHERE short_url = $1`
	row := r.DB.QueryRow(query, shortURL)

	var url model.URL
	err := row.Scan(&url.ID, &url.ShortURL, &url.OriginalURL, &url.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) ShortURLExists(shortURL string) (bool, error) {
	query := `SELECT 1 FROM urls WHERE short_url = $1`
	row := r.DB.QueryRow(query, shortURL)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

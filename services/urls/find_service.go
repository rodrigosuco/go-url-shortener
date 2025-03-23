package services

import (
	"context"
	"errors"

	"github.com/rodrigosuco/go-url-shortener/internal/database"
	"github.com/rodrigosuco/go-url-shortener/models"
)

func FindOriginalURl(short_url string) (*models.Url, error) {
	query := `SELECT id, original_url, short_url, created_at FROM urls WHERE short_url = $1`
	rows, err := database.DB.Query(context.Background(), query, short_url)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var url models.Url
	if rows.Next() {
		if err := rows.Scan(&url.ID, &url.OriginalUrl, &url.ShortUrl, &url.CreatedAt); err != nil {
			return nil, err
		}
		return &url, nil
	}

	return nil, errors.New("URL not found")
}

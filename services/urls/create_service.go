package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rodrigosuco/go-url-shortener/internal/database"
	"github.com/rodrigosuco/go-url-shortener/models"
)

func CreateUrl(url models.Url) (models.Url, error)  {

	url.ShortUrl = randomString(5)

	url, err := saveUrl(url)

	if err != nil {
		return url, err
	}

	return url, nil

}

func saveUrl(url models.Url) (models.Url, error) {
	query := `INSERT INTO urls (original_url, short_url)
						VALUES ($1, $2)
						RETURNING id, original_url, short_url, created_at;`

	err := database.DB.QueryRow(
	context.Background(),
	query,
	url.OriginalUrl,
	url.ShortUrl,
	).Scan(
	&url.ID,
	&url.OriginalUrl,
	&url.ShortUrl,
	&url.CreatedAt,
	)

	if err != nil {
		return models.Url{}, fmt.Errorf("failed to create url: %v", err)
	}
	return url, nil
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}

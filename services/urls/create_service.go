package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/rodrigosuco/go-url-shortener/internal/database"
	"github.com/rodrigosuco/go-url-shortener/models"
)

const shortUrlLength = 5
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func CreateUrl(url models.Url) (models.Url, error) {
	var err error
	for {
		url.ShortUrl = randomString(shortUrlLength)
		exists, _ := FindOriginalURl(url.ShortUrl)
		if exists == nil {
			break
		}
	}

	url, err = saveUrl(url)
	if err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func saveUrl(url models.Url) (models.Url, error) {
	query := `INSERT INTO urls (original_url, short_url)
	          VALUES ($1, $2)
	          RETURNING id, original_url, short_url, created_at`

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
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

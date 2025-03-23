package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigosuco/go-url-shortener/models"
	"github.com/rodrigosuco/go-url-shortener/services/urls"
)
func CreateShortUrl(ctx *gin.Context)  {
	var url models.Url

	body, _ := io.ReadAll(ctx.Request.Body)

	validateBody(body, ctx)

	ctx.ShouldBindBodyWithJSON(&url)

	shornetedUrl, err := services.CreateUrl(url)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating url, error: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"url": shornetedUrl,
	})
}

func validateBody(body []byte, ctx *gin.Context)  {
	var requestData map[string]any
	if err := json.Unmarshal(body, &requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	originalUrl, exists := requestData["original_url"]
	if !exists || originalUrl == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing original_url param"})
		return
	}
}
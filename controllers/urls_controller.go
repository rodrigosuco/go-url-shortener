package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigosuco/go-url-shortener/models"
	"github.com/rodrigosuco/go-url-shortener/services/urls"
)

func GetOriginalUrl(ctx *gin.Context) {
	short_url := ctx.Param("short_url")

	url, err := services.FindOriginalURl(short_url)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"url": url})
}

func CreateShortUrl(ctx *gin.Context)  {
	var url models.Url

	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if url.OriginalUrl == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing original_url param"})
		return
	}

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

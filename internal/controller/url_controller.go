package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sancheschris/shorty/internal/service"
)

type ShortenRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	url, err := service.CreateShortURL(req.OriginalURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a short URL"})
		return
	}

	shortURL := c.Request.Host + "/" + url.ShortCode
	c.JSON(http.StatusOK, ShortenResponse{ShortURL: "http://" + shortURL})
}

func RedirectURL(c *gin.Context) {
	code := c.Param("shortCode")

	url, err := service.GetOriginalURL(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
		return
	}

	c.Redirect(http.StatusFound, url.OriginalURL)
}
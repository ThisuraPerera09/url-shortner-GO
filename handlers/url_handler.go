package handlers

import (
	"net/http"
	"strconv"
	"url-shortener/models"
	"url-shortener/service"
	"url-shortener/storage"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service *service.URLService
	baseURL string
}

func NewURLHandler(service *service.URLService, baseURL string) *URLHandler {
	return &URLHandler{
		service: service,
		baseURL: baseURL,
	}
}

// ShortenURL handles POST /api/shorten
func (h *URLHandler) ShortenURL(c *gin.Context) {
	var req models.ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	url, err := h.service.ShortenURL(req.URL, req.CustomCode)
	if err != nil {
		if err == storage.ErrAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Custom code already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}

	response := models.ShortenResponse{
		ShortCode:   url.ShortCode,
		ShortURL:    h.baseURL + "/" + url.ShortCode,
		OriginalURL: url.OriginalURL,
	}

	c.JSON(http.StatusCreated, response)
}

// RedirectURL handles GET /:shortCode
func (h *URLHandler) RedirectURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	url, err := h.service.GetURL(shortCode)
	if err != nil {
		if err == storage.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
}

// GetStats handles GET /api/stats/:shortCode
func (h *URLHandler) GetStats(c *gin.Context) {
	shortCode := c.Param("shortCode")

	url, err := h.service.GetStats(shortCode)
	if err != nil {
		if err == storage.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stats"})
		return
	}

	response := models.StatsResponse{
		ShortCode:    url.ShortCode,
		OriginalURL:  url.OriginalURL,
		Clicks:       url.Clicks,
		CreatedAt:    url.CreatedAt,
		LastAccessed: url.LastAccessed,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteURL handles DELETE /api/urls/:shortCode
func (h *URLHandler) DeleteURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	err := h.service.DeleteURL(shortCode)
	if err != nil {
		if err == storage.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL deleted successfully"})
}

// ListURLs handles GET /api/urls
func (h *URLHandler) ListURLs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	urls, err := h.service.ListURLs(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URLs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"urls":   urls,
		"limit":  limit,
		"offset": offset,
		"count":  len(urls),
	})
}

// HealthCheck handles GET /api/health
func (h *URLHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "url-shortener",
	})
}

package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services"
)

// MediaHandlers handles media route
type MediaHandlers struct {
	media services.MediaService
}

// NewMediaHandlers creates a media handlers instance
func NewMediaHandlers(DB *gorm.DB) MediaHandlers {
	return MediaHandlers{
		media: services.NewMediaService(DB),
	}
}

// Search search media corresponding to given query parameters
func (h *MediaHandlers) Search(c *gin.Context) {
	inputs := db.SearchMediaInputs{
		LibraryID:  c.Query("libraryId"),
		SerieID:    c.Query("serieId"),
		MediaType:  c.Query("mediaType"),
		MediaIndex: c.Query("mediaIndex"),
	}

	if inputs.LibraryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "LibraryId is mandatory"})
		return
	}

	medias, err := h.media.Search(inputs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error searching for media"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medias})
}

// Get get media info
func (h *MediaHandlers) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	media, err := h.media.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	c.JSON(http.StatusOK, media)
}

// GetContent return the media content
func (h *MediaHandlers) GetContent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageIndex = 0
	}

	err = h.media.StreamMediaPage(uint(id), pageIndex, c.Writer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	c.Header("Content-type", "image/jpg")
	c.Header("Content-Disposition", "inline")
	c.Writer.WriteHeader(http.StatusOK)
}

// MarkAsRead mark media as read
func (h *MediaHandlers) MarkAsRead(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	err = h.media.MarkAsRead(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// MarkAsUnread mark media as read
func (h *MediaHandlers) MarkAsUnread(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	err = h.media.MarkAsUnread(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

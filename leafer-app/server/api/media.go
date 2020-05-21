package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/utils"
)

// SearchMedia search media corresponding to given query parameters
func SearchMedia(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)
	store := db.NewMediaStore(conn)
	inputs := db.SearchMediaInputs{
		LibraryID:     c.Query("libraryId"),
		ParentMediaID: c.Query("parentMediaId"),
		MediaType:     c.Query("mediaType"),
		MediaIndex:    c.Query("mediaIndex"),
	}

	if inputs.LibraryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "LibraryId is mandatory"})
		return
	}

	medias, err := store.Search(inputs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error searching for media"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medias})
}

// FindMedia get media info
func FindMedia(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)
	store := db.NewMediaStore(conn)

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	media, err := store.GetMedia(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	c.JSON(http.StatusOK, media)
}

// GetMediaContent return the media content
func GetMediaContent(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)
	store := db.NewMediaStore(conn)

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	media, err := store.GetMedia(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageIndex = 0
	}

	err = utils.StreamImageFromZip(media.Path, pageIndex, c.Writer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	c.Header("Content-type", "image/jpg")
	c.Header("Content-Disposition", "inline")
	c.Writer.WriteHeader(http.StatusOK)
}

// MarkMediaAsRead mark media as read
func MarkMediaAsRead(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)
	store := db.NewMediaStore(conn)

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	media, err := store.GetMedia(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	viewedAt := time.Now()
	err = store.UpdateLastViewed(uint(id), &viewedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, media)
}

// MarkMediaAsUnread mark media as read
func MarkMediaAsUnread(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)
	store := db.NewMediaStore(conn)

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	media, err := store.GetMedia(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	err = store.UpdateLastViewed(uint(id), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, media)
}

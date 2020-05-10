package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/utils"
)

// GetMedia get media info dependening on its type
func GetMedia(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var id = c.Param("id")
	var media db.Media
	var queryMedia = conn.First(&media, id)
	if queryMedia.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	if media.Type == "COLLECTION" {
		var medias []db.Media
		conn.Where("parent_media_id = ?", id).Find(&medias)
		media.Medias = &medias
	}

	c.JSON(http.StatusOK, gin.H{"data": media})
}

// GetMediaContent return the media content
func GetMediaContent(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var id = c.Param("id")
	var media db.Media
	var queryMedia = conn.First(&media, id)
	if queryMedia.Error != nil || media.Type == "COLLECTION" {
		c.JSON(http.StatusNotFound, gin.H{"error": "media content not found"})
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		pageIndex = 0
	}

	err = utils.StreamImageFromZip(media.FilePath, pageIndex, c.Writer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	c.Header("Content-type", "image/jpg")
	c.Header("Content-Disposition", "inline")
	c.Writer.WriteHeader(http.StatusOK)
}

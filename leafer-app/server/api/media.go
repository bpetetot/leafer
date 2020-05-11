package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/utils"
)

// SearchMedia search media corresponding to given query parameters
// valid query parameters are:
// - libraryId
// - parentMediaId
// - mediaType
// - mediaIndex
func SearchMedia(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var libraryID = c.Query("libraryId")
	if libraryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "LibraryId is mandatory"})
		return
	}
	conn = conn.Where("library_id = ?", libraryID)

	var parentMediaID = c.Query("parentMediaId")
	if parentMediaID != "" {
		conn = conn.Where("parent_media_id = ?", parentMediaID)
	}

	var mediaType = c.Query("mediaType")
	if mediaType != "" {
		conn = conn.Where("type = ?", mediaType)
	}

	var mediaIndex = c.Query("mediaIndex")
	if mediaIndex != "" {
		conn = conn.Where("media_index = ?", mediaIndex)
	}

	var medias []db.Media
	var queryMedia = conn.Find(&medias)
	if queryMedia.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "error searching for media"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medias})
}

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
		conn.Where("parent_media_id = ?", id).Order("volume").Find(&medias)
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

	err = utils.StreamImageFromZip(media.Path, pageIndex, c.Writer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
		return
	}

	c.Header("Content-type", "image/jpg")
	c.Header("Content-Disposition", "inline")
	c.Writer.WriteHeader(http.StatusOK)
}

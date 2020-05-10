package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
)

// ListMedia list media files of a collection
func ListMedia(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var id = c.Param("id")
	var media db.Media
	var queryMedia = conn.First(&media, id)

	if queryMedia.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "media not found"})
		return
	}

	if media.ParentMediaID == 0 {
		var medias []db.Media
		conn.Where("parent_media_id = ?", id).Find(&medias)
		media.Medias = &medias
	}

	c.JSON(http.StatusOK, gin.H{"data": media})
}

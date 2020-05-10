package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
)

// CreateLibraryInput struct to add a new library with validation
type CreateLibraryInput struct {
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

// UpdateLibraryInput struct to update a library
type UpdateLibraryInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ListLibraries GET /libraries
func ListLibraries(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var libraries []db.Library
	conn.Find(&libraries)

	c.JSON(http.StatusOK, gin.H{"data": libraries})
}

// FindLibrary GET /libraries/:id
func FindLibrary(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var library db.Library
	var query = conn.Where("id = ?", c.Param("id")).First(&library)

	if query.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "library not found"})
		return
	}

	var medias []db.Media
	var id = c.Param("id")
	conn.Where("library_id = ? AND parent_media_id = ?", id, 0).Find(&medias)

	library.Medias = &medias

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// CreateLibrary POST /libraries
func CreateLibrary(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var input CreateLibraryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	library := db.Library{Name: input.Name, Path: input.Path}
	conn.Create(&library)

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// UpdateLibrary update the given library
func UpdateLibrary(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var library db.Library
	var query = conn.Where("id = ?", c.Param("id")).First(&library)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateLibraryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn.Model(&library).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// DeleteLibrary delete the given library
func DeleteLibrary(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var library db.Library
	var query = conn.Where("id = ?", c.Param("id")).First(&library)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.DeleteLibrary(&library, conn)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

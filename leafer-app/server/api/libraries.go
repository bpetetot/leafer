package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/scanners"
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

	c.JSON(http.StatusOK, libraries)
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

	c.JSON(http.StatusOK, library)
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

	c.JSON(http.StatusOK, library)
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

	c.JSON(http.StatusOK, library)
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

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// ScanLibraryAsync scan files for the given library
func ScanLibraryAsync(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	var library db.Library
	var query = conn.Where("id = ?", c.Param("id")).First(&library)

	if query.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "library not found"})
		return
	}

	go func(library *db.Library, conn *gorm.DB) {
		scanners.ScanLibrary(library, conn)
		scanners.ScanMedias(library, conn)
	}(&library, conn)

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

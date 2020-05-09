package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	models "github.com/bpetetot/leafer/db"
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
	db := c.MustGet("db").(*gorm.DB)

	var libraries []models.Library
	db.Find(&libraries)

	c.JSON(http.StatusOK, gin.H{"data": libraries})
}

// CreateLibrary POST /libraries
func CreateLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateLibraryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	library := models.Library{Name: input.Name, Path: input.Path}
	db.Create(&library)

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// FindLibrary GET /libraries/:id
func FindLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var library models.Library
	var query = db.Where("id = ?", c.Param("id")).First(&library)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// UpdateLibrary update the given library
func UpdateLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var library models.Library
	var query = db.Where("id = ?", c.Param("id")).First(&library)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateLibraryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&library).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": library})
}

// DeleteLibrary delete the given library
func DeleteLibrary(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var library models.Library
	var query = db.Where("id = ?", c.Param("id")).First(&library)

	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&library)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

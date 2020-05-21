package api

import (
	"net/http"
	"strconv"

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

// ListLibraries GET /libraries
func ListLibraries(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)
	libraries, err := db.FindLibraries(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, libraries)
}

// FindLibrary GET /libraries/:id
func FindLibrary(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	library, err := db.GetLibrary(conn, uint(id))
	if err != nil {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	library, err := db.CreateLibrary(conn, input.Name, input.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, library)
}

// DeleteLibrary delete the given library
func DeleteLibrary(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	_, err = db.GetLibrary(conn, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "library not found"})
		return
	}

	err = db.DeleteLibrary(conn, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// ScanLibraryAsync scan files for the given library
func ScanLibraryAsync(c *gin.Context) {
	conn := c.MustGet("db").(*gorm.DB)

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	library, err := db.GetLibrary(conn, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "library not found"})
		return
	}

	go func(library *db.Library, conn *gorm.DB) {
		scanners.ScanLibrary(library, conn)
		scanners.ScanMedias(library, conn)
	}(library, conn)

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

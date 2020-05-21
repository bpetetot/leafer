package server

import (
	"net/http"
	"strconv"

	"github.com/bpetetot/leafer/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// LibraryHandlers handles libraries route
type LibraryHandlers struct {
	library services.LibraryService
	scanner services.ScannerService
	scraper services.ScraperService
}

// NewLibraryHandlers creates a library handlers instance
func NewLibraryHandlers(DB *gorm.DB) LibraryHandlers {
	return LibraryHandlers{
		library: services.NewLibraryService(DB),
		scanner: services.NewScannerService(DB),
		scraper: services.NewScraperService(DB),
	}
}

// Find GET /libraries
func (h *LibraryHandlers) Find(c *gin.Context) {
	libraries, err := h.library.Find()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, libraries)
}

// Get GET /libraries/:id
func (h *LibraryHandlers) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	lib, err := h.library.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "library not found"})
		return
	}

	c.JSON(http.StatusOK, lib)
}

// CreateLibraryInput struct to add a new library with validation
type CreateLibraryInput struct {
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

// Create POST /libraries
func (h *LibraryHandlers) Create(c *gin.Context) {
	var input CreateLibraryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	lib, err := h.library.Create(input.Name, input.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, lib)
}

// Delete deletes the given library
func (h *LibraryHandlers) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	err = h.library.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

// Scan scan files for the given library
func (h *LibraryHandlers) Scan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not a valid integer"})
		return
	}

	go func(id uint) {
		h.scanner.ScanLibrary(id)
		h.scraper.ScrapLibrary(id)
	}(uint(id))

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

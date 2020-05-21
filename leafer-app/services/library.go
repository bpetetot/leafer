package services

import (
	"github.com/jinzhu/gorm"

	"github.com/bpetetot/leafer/db"
)

// LibraryService gives access to library services
type LibraryService struct {
	libraryStore db.LibraryStore
	mediaStore   db.MediaStore
}

// NewLibraryService creates a library service instance
func NewLibraryService(DB *gorm.DB) LibraryService {
	return LibraryService{
		libraryStore: db.NewLibraryStore(DB),
		mediaStore:   db.NewMediaStore(DB),
	}
}

// Find returns all libraries
func (s *LibraryService) Find() (*[]db.Library, error) {
	return s.libraryStore.Find()
}

// Get returns the library for the given id
func (s *LibraryService) Get(id uint) (*db.Library, error) {
	return s.libraryStore.Get(id)
}

// Create adds a new library in database
func (s *LibraryService) Create(name string, path string) (*db.Library, error) {
	return s.libraryStore.Create(name, path)
}

// Delete deletes the library for the given id
func (s *LibraryService) Delete(id uint) error {
	err := s.mediaStore.DeleteMediasLibrary(id)
	if err != nil {
		return err
	}

	err = s.libraryStore.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

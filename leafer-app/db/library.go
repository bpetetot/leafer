package db

import (
	"github.com/jinzhu/gorm"
)

// LibraryStore is the store giving access and writes to libraries
type LibraryStore struct {
	DB *gorm.DB
}

// NewLibraryStore creates a library store instance
func NewLibraryStore(db *gorm.DB) LibraryStore {
	return LibraryStore{DB: db}
}

// FindLibraries returns all libraries in the db
func (s LibraryStore) FindLibraries() ([]Library, error) {
	var libraries []Library
	query := s.DB.Find(&libraries)
	if query.Error != nil {
		return nil, query.Error
	}

	return libraries, query.Error
}

// GetLibrary returns the library with the given id
func (s LibraryStore) GetLibrary(id uint) (*Library, error) {
	var library Library
	query := s.DB.First(&library, id)
	if query.Error != nil {
		return nil, query.Error
	}

	return &library, query.Error
}

// CreateLibrary adds new library in db and returns it
func (s LibraryStore) CreateLibrary(name string, path string) (*Library, error) {
	library := Library{Name: name, Path: path}
	query := s.DB.Create(&library)
	if query.Error != nil {
		return nil, query.Error
	}

	query = s.DB.Last(&library)
	if query.Error != nil {
		return nil, query.Error
	}

	return &library, nil
}

// DeleteLibrary delete the library and its medias
func (s LibraryStore) DeleteLibrary(id uint) error {
	err := DeleteMediasLibrary(s.DB, id)
	if err != nil {
		return err
	}

	query := s.DB.Unscoped().Delete(Library{ID: id})
	return query.Error
}

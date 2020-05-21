package db

import (
	"github.com/jinzhu/gorm"
)

// LibraryStore is a interface for library store
type LibraryStore interface {
	Find() (*[]Library, error)
	Get(id uint) (*Library, error)
	Create(name string, path string) (*Library, error)
	Delete(id uint) error
}

type libraryRepo struct {
	DB *gorm.DB
}

// NewLibraryStore creates a library store instance
func NewLibraryStore(db *gorm.DB) LibraryStore {
	return &libraryRepo{DB: db}
}

// Find returns all libraries in the db
func (r *libraryRepo) Find() (*[]Library, error) {
	var libraries []Library
	query := r.DB.Find(&libraries)
	if query.Error != nil {
		return nil, query.Error
	}

	return &libraries, query.Error
}

// Get returns the library with the given id
func (r *libraryRepo) Get(id uint) (*Library, error) {
	var library Library
	query := r.DB.First(&library, id)
	if query.Error != nil {
		return nil, query.Error
	}

	return &library, query.Error
}

// Create adds new library in db and returns it
func (r *libraryRepo) Create(name string, path string) (*Library, error) {
	library := Library{Name: name, Path: path}
	query := r.DB.Create(&library)
	if query.Error != nil {
		return nil, query.Error
	}

	query = r.DB.Last(&library)
	if query.Error != nil {
		return nil, query.Error
	}

	return &library, nil
}

// Delete deletes the library
func (r *libraryRepo) Delete(id uint) error {
	query := r.DB.Unscoped().Delete(Library{ID: id})
	return query.Error
}

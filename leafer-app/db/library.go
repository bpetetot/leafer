package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Library model
type Library struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Medias    *[]Media  `json:"medias,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// FindLibraries returns all libraries in the db
func FindLibraries(db *gorm.DB) ([]Library, error) {
	var libraries []Library
	query := db.Find(&libraries)
	if query.Error != nil {
		return nil, query.Error
	}

	return libraries, query.Error
}

// GetLibrary returns the library with the given id
func GetLibrary(db *gorm.DB, id uint) (*Library, error) {
	var library Library
	query := db.First(&library, id)
	if query.Error != nil {
		return nil, query.Error
	}

	return &library, query.Error
}

// CreateLibrary adds new library in db and returns it
func CreateLibrary(db *gorm.DB, name string, path string) (*Library, error) {
	library := Library{Name: name, Path: path}
	query := db.Create(&library)
	if query.Error != nil {
		return nil, query.Error
	}

	query = db.Last(&library)
	if query.Error != nil {
		return nil, query.Error
	}

	return &library, nil
}

// DeleteLibrary delete the library and its medias
func DeleteLibrary(db *gorm.DB, id uint) error {
	err := DeleteMediasLibrary(db, id)
	if err != nil {
		return err
	}

	query := db.Unscoped().Delete(Library{ID: id})
	return query.Error
}

package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// MediaStore is the store giving access and writes to medias
type MediaStore struct {
	DB *gorm.DB
}

// NewMediaStore creates a media store instance
func NewMediaStore(db *gorm.DB) MediaStore {
	return MediaStore{DB: db}
}

// GetMedia get media info dependening on its type
func (s MediaStore) GetMedia(id uint) (*Media, error) {
	var media Media
	query := s.DB.First(&media, id)
	if query.Error != nil {
		return nil, query.Error
	}
	return &media, nil
}

// SearchMediaInputs represents possible search inputs
type SearchMediaInputs struct {
	LibraryID     string
	ParentMediaID string
	MediaType     string
	MediaIndex    string
}

// Search search medias corresponding to given search inputs
func (s MediaStore) Search(inputs SearchMediaInputs) ([]Media, error) {
	query := s.DB.Where("library_id = ?", inputs.LibraryID)

	if inputs.ParentMediaID != "" {
		query = query.Where("parent_media_id = ?", inputs.ParentMediaID)
	}
	if inputs.MediaType != "" {
		query = query.Where("type = ?", inputs.MediaType)
	}
	if inputs.MediaIndex != "" {
		query = query.Where("media_index = ?", inputs.MediaIndex)
	}

	var medias []Media
	query = query.Order("mediaIndex").Find(&medias)
	if query.Error != nil {
		return nil, query.Error
	}
	return medias, nil
}

// UpdateLastViewed sets the last viewed value of the media
func (s MediaStore) UpdateLastViewed(id uint, when *time.Time) error {
	query := s.DB.Model(Media{ID: id}).Update("LastViewedAt", when)
	return query.Error
}

// DeleteMediasLibrary deletes library's media
func DeleteMediasLibrary(db *gorm.DB, LibraryID uint) error {
	query := db.Unscoped().Where("library_id = ?", LibraryID).Delete(&Media{})
	return query.Error
}

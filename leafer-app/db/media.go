package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// MediaStore is the store giving access and writes to medias
type MediaStore interface {
	Get(id uint) (*Media, error)
	Search(inputs SearchMediaInputs) (*[]Media, error)
	CountSearch(inputs SearchMediaInputs) int
	Create(media *Media) (*Media, error)
	Update(id uint, media *Media) error
	UpdateLastViewed(id uint, when *time.Time) error
	Delete(id uint) error
	DeleteMediasLibrary(id uint) error
}

type mediaRepo struct {
	DB *gorm.DB
}

// NewMediaStore creates a media store instance
func NewMediaStore(db *gorm.DB) MediaStore {
	return &mediaRepo{DB: db}
}

// Get get media info dependening on its type
func (r *mediaRepo) Get(id uint) (*Media, error) {
	var media Media
	query := r.DB.First(&media, id)
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

func buildSearchQuery(db *gorm.DB, inputs SearchMediaInputs) *gorm.DB {
	query := db.Where("library_id = ?", inputs.LibraryID)

	if inputs.ParentMediaID != "" {
		query = query.Where("parent_media_id = ?", inputs.ParentMediaID)
	}
	if inputs.MediaType != "" {
		query = query.Where("type = ?", inputs.MediaType)
	}
	if inputs.MediaIndex != "" {
		query = query.Where("media_index = ?", inputs.MediaIndex)
	}
	return query
}

// Search search medias corresponding to given search inputs
func (r *mediaRepo) Search(inputs SearchMediaInputs) (*[]Media, error) {
	query := buildSearchQuery(r.DB, inputs)

	var medias []Media
	query = query.Order("mediaIndex").Find(&medias)
	if query.Error != nil {
		return nil, query.Error
	}
	return &medias, nil
}

// CountSearch count medias corresponding to given search inputs
func (r *mediaRepo) CountSearch(inputs SearchMediaInputs) int {
	query := buildSearchQuery(r.DB.Model(Media{}), inputs)

	var count int
	query = query.Count(&count)
	if query.Error != nil {
		return 0
	}
	return count
}

// Create creates the media
func (r *mediaRepo) Create(media *Media) (*Media, error) {
	query := r.DB.Create(&media)
	if query.Error != nil {
		return nil, query.Error
	}

	query = r.DB.Last(&media)
	if query.Error != nil {
		return nil, query.Error
	}
	return media, nil
}

// Update updates the media
func (r *mediaRepo) Update(id uint, media *Media) error {
	query := r.DB.Model(Media{ID: id}).Update(media)
	return query.Error
}

// UpdateLastViewed sets the last viewed value of the media
func (r *mediaRepo) UpdateLastViewed(id uint, when *time.Time) error {
	query := r.DB.Model(Media{ID: id}).Update("LastViewedAt", when)
	return query.Error
}

// Delete deletes the media
func (r *mediaRepo) Delete(id uint) error {
	query := r.DB.Unscoped().Where("ID = ?", id).Delete(&Media{})
	return query.Error
}

// DeleteMediasLibrary deletes library's media
func (r *mediaRepo) DeleteMediasLibrary(id uint) error {
	query := r.DB.Unscoped().Where("library_id = ?", id).Delete(&Media{})
	return query.Error
}

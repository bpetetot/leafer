package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// MediaStore is the store giving access and writes to medias
type MediaStore interface {
	Get(id uint) (*Media, error)
	Search(inputs SearchMediaInputs) (*[]Media, error)
	CountSearch(inputs SearchMediaInputs) int
	GetFirstMediaSerie(libraryID uint, serieID uint) (*Media, error)
	Create(media *Media) (*Media, error)
	Update(id uint, media *Media) error
	UpdateLastViewed(id uint, when *time.Time) error
	Delete(id uint) error
	UpdateMediasLibraryScanningStatus(id uint, scanningStatus string) error
	DeleteMediasLibrary(id uint) error
	DeleteMediasWithScanningStatus(libraryID uint) error
	ExistMediaPath(libraryID uint, path string) (*Media, error)
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
	LibraryID  string
	SerieID    string
	MediaType  string
	MediaIndex string
	Path       string
}

func buildSearchQuery(db *gorm.DB, inputs SearchMediaInputs) *gorm.DB {
	query := db.Where("library_id = ?", inputs.LibraryID)

	if inputs.SerieID != "" {
		query = query.Where("serie_id = ?", inputs.SerieID)
	}
	if inputs.MediaType != "" {
		query = query.Where("type = ?", inputs.MediaType)
	}
	if inputs.MediaIndex != "" {
		query = query.Where("media_index = ?", inputs.MediaIndex)
	}
	if inputs.Path != "" {
		query = query.Where("path = ?", inputs.Path)
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

// GetFirstMediaSerie get the first media of a serie
func (r *mediaRepo) GetFirstMediaSerie(libraryID uint, serieID uint) (*Media, error) {
	query := buildSearchQuery(r.DB.Model(Media{}), SearchMediaInputs{LibraryID: fmt.Sprint(libraryID), SerieID: fmt.Sprint(serieID)})

	var media Media
	query = query.First(&media)
	if query.Error != nil {
		return nil, query.Error
	}
	return &media, nil
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

// DeleteMediasWithScanningStatus deletes library's media with "scanning" status
func (r *mediaRepo) DeleteMediasWithScanningStatus(id uint) error {
	query := r.DB.Unscoped().Where("library_id = ?", id).Where("scanning_status = ?", "scanning").Delete(&Media{})
	return query.Error
}

// Update update all medias library scanning status
func (r *mediaRepo) UpdateMediasLibraryScanningStatus(id uint, scanningStatus string) error {
	query := r.DB.Model(&Media{}).Where("library_id = ?", id).Update(&Media{ScanningStatus: scanningStatus})
	return query.Error
}

// ExistMediaPath returns the medias if one exists for the library and path
func (r *mediaRepo) ExistMediaPath(libraryID uint, path string) (*Media, error) {
	medias, err := r.Search(SearchMediaInputs{LibraryID: fmt.Sprint(libraryID), Path: path})
	if err != nil {
		return nil, err
	}
	if len(*medias) != 1 {
		return nil, nil
	}
	return &(*medias)[0], nil
}

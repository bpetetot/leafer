package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Media model
type Media struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`

	LibraryID uint     `json:"-"`
	Library   *Library `json:"-"`

	ParentMediaID uint   `json:"-"`
	ParentMedia   *Media `json:"-"`

	Title       string `json:"title,omitempty"`
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	Country     string `json:"countryOfOrigin,omitempty"`
	CoverImage  string `json:"coverImage,omitempty"`
	BannerImage string `json:"bannerImage,omitempty"`
	Score       int    `json:"averageScore,omitempty"`

	EstimatedName string `json:"estimatedName,omitempty"`
	Volume        int    `json:"volume,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	FileExtension string `json:"fileExtention,omitempty"`
	PageCount     int    `json:"pageCount,omitempty"`
	MediaIndex    int    `json:"mediaIndex,omitempty"`

	MediaCount int `json:"mediaCount,omitempty"`

	LastViewedAt *time.Time `json:"lastViewedAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// DeleteMediasLibrary deletes library's media
func DeleteMediasLibrary(db *gorm.DB, LibraryID uint) error {
	query := db.Unscoped().Where("LibraryID = ?", LibraryID).Delete(&Media{})
	return query.Error
}

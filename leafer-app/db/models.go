package db

import "time"

// Library model
type Library struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	Name           string    `json:"name"`
	Path           string    `json:"path"`
	Medias         *[]Media  `json:"medias,omitempty"`
	ScanningStatus string    `json:"scanningStatus"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// Media model
type Media struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Type string `json:"type,omitempty"`
	Path string `json:"path,omitempty"`

	LibraryID uint     `json:"-"`
	Library   *Library `json:"-"`

	SerieID uint   `json:"-"`
	Serie   *Media `json:"-"`

	Title           string `json:"title,omitempty"`
	Category        string `json:"category,omitempty"`
	Description     string `json:"description,omitempty"`
	Country         string `json:"countryOfOrigin,omitempty"`
	CoverImageURL   string `json:"coverImageUrl,omitempty"`
	CoverImageLocal string `json:"coverImageLocal,omitempty"`
	Volume          int    `json:"volume,omitempty"`
	Score           int    `json:"averageScore,omitempty"`
	PageCount       int    `json:"pageCount,omitempty"`

	FileName   string `json:"fileName,omitempty"`
	MediaIndex int    `json:"mediaIndex,omitempty"`

	MediaCount int `json:"mediaCount,omitempty"`

	ScanningStatus string     `json:"scanningStatus"`
	LastViewedAt   *time.Time `json:"lastViewedAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

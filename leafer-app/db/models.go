package db

import "time"

// Library model
type Library struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Medias    *[]Media  `json:"medias,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

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

	Volume     int    `json:"volume,omitempty"`
	FileName   string `json:"fileName,omitempty"`
	PageCount  int    `json:"pageCount,omitempty"`
	MediaIndex int    `json:"mediaIndex,omitempty"`

	MediaCount int `json:"mediaCount,omitempty"`

	LastViewedAt *time.Time `json:"lastViewedAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

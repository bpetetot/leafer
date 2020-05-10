package db

import (
	"time"
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

// Media model
type Media struct {
	ID uint `json:"id" gorm:"primary_key"`

	LibraryID uint     `json:"-"`
	Library   *Library `json:"-"`

	ParentMediaID uint   `json:"-"`
	ParentMedia   *Media `json:"-"`

	Title       string `json:"title,omitempty"`
	TitleNative string `json:"titleNative,omitempty"`
	Type        string `json:"type,omitempty"`
	Volume      int    `json:"volume,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	Country     string `json:"countryOfOrigin,omitempty"`
	CoverImage  string `json:"coverImage,omitempty"`
	BannerImage string `json:"bannerImage,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Score       int    `json:"averageScore,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
	EndDate     string `json:"endDate,omitempty"`

	EstimatedName string `json:"estimatedName,omitempty"`
	FilePath      string `json:"filePath,omitempty"`
	FileDir       string `json:"fileDir,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	FileExtension string `json:"fileExtention,omitempty"`

	Medias *[]Media `json:"medias,omitempty" gorm:"-"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

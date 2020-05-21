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

// DeleteLibrary delete the library and its medias
func DeleteLibrary(library *Library, conn *gorm.DB) {
	DeleteLibraryContent(library, conn)
	conn.Unscoped().Delete(&library)
}

// DeleteLibraryContent delete library's media
func DeleteLibraryContent(library *Library, conn *gorm.DB) {
	var medias []Media
	conn.Model(library).Association("Medias").Find(&medias)

	for _, media := range medias {
		conn.Unscoped().Delete(&media)
	}
}

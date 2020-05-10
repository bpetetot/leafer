package db

import (
	"github.com/jinzhu/gorm"
)

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

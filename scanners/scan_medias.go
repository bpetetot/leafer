package scanners

import (
	"log"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/scrapers"
	"github.com/jinzhu/gorm"
)

// ScanMedias scans medias from database to get metadata info
func ScanMedias(library *db.Library, conn *gorm.DB) {
	var medias []db.Media
	conn.Model(&library).Association("Medias").Find(&medias)

	for _, media := range medias {
		if media.ParentMediaID == 0 {
			log.Printf("Scan media collection for %v", media.EstimatedName)
			found := scrapers.Scrap(media.EstimatedName)
			conn.Model(&media).Updates(found)
		}
	}
}

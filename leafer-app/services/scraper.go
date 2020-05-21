package services

import (
	"fmt"
	"log"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/scrapers"
	"github.com/jinzhu/gorm"
)

// ScraperService exposes service to scan a library
type ScraperService struct {
	library db.LibraryStore
	media   db.MediaStore
}

// NewScraperService creates a library scanner service instance
func NewScraperService(DB *gorm.DB) ScraperService {
	return ScraperService{
		library: db.NewLibraryStore(DB),
		media:   db.NewMediaStore(DB),
	}
}

// ScrapLibrary scans the given library id
func (s *ScraperService) ScrapLibrary(id uint) error {
	log.Printf("Scan media for library [%v]", id)

	medias, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(id)})
	if err != nil {
		return err
	}

	for _, media := range *medias {
		if media.ParentMediaID == 0 {
			log.Printf("Scan media collection for %v", media.EstimatedName)
			found := scrapers.Scrap(media.EstimatedName)
			found.MediaCount = s.media.CountSearch(db.SearchMediaInputs{LibraryID: fmt.Sprint(id), ParentMediaID: fmt.Sprint(media.ID)})
			s.media.Update(media.ID, &found)
		}
	}
	return nil
}

package services

import (
	"fmt"
	"log"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/scrapers"
	"github.com/jinzhu/gorm"
)

// MetadataService exposes service to scan a library
type MetadataService struct {
	library db.LibraryStore
	media   db.MediaStore
}

// NewMetadataService creates a library scanner service instance
func NewMetadataService(DB *gorm.DB) MetadataService {
	return MetadataService{
		library: db.NewLibraryStore(DB),
		media:   db.NewMediaStore(DB),
	}
}

// ScanLibrary scans the given library id
func (s *MetadataService) ScanLibrary(id uint) error {
	log.Printf("Scan media for library [%v]", id)

	collections, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(id), ParentMediaID: "0"})
	if err != nil {
		return err
	}

	for _, media := range *collections {
		log.Printf("Scan media collection for %v", media.EstimatedName)
		found := scrapers.Scrap(media.EstimatedName)
		found.MediaCount = s.media.CountSearch(db.SearchMediaInputs{LibraryID: fmt.Sprint(id), ParentMediaID: fmt.Sprint(media.ID)})
		s.media.Update(media.ID, &found)
	}
	return nil
}

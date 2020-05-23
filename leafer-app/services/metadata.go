package services

import (
	"fmt"
	"log"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/comicfile"
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
	log.Printf("Scan metadata library [%v]", id)

	collections, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(id), ParentMediaID: "0"})
	if err != nil {
		return err
	}

	for _, collection := range *collections {
		log.Printf("Metadata collection [%v]", collection.Path)

		medias, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(id), ParentMediaID: fmt.Sprint(collection.ID)})
		if err != nil {
			return err
		}

		for _, media := range *medias {
			comic, err := comicfile.New(media.Path)
			if err != nil {
				s.media.Delete(media.ID)
				continue
			}
			metadata, err := comic.ExtractMetadata()
			if err != nil {
				continue
			}

			// Media metadata
			media.PageCount = metadata.PageCount
			media.Volume = metadata.Volume
			s.media.Update(media.ID, &media)
		}

		// Collection metadata
		collectionScrapped := scrapers.Scrap(collection.FileName)
		collectionScrapped.MediaCount = len(*medias)
		s.media.Update(collection.ID, &collectionScrapped)
	}
	return nil
}

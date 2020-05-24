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
		s.ScanCollection(collection)
	}
	return nil
}

// ScanCollection scans metadata for the whole collection and its medias
func (s *MetadataService) ScanCollection(collection db.Media) error {
	log.Printf("Metadata collection [%v]", collection.Path)

	medias, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(collection.LibraryID), ParentMediaID: fmt.Sprint(collection.ID)})
	if err != nil {
		return err
	}
	mediaCount := len(*medias)

	// Get collection metadata from the first media metadata
	collectionUpdated := false
	if mediaCount > 0 {
		comic, _ := comicfile.New((*medias)[0].Path)
		collectionMetadata, _ := comic.ExtractMetadata()
		collectionMetadata.MediaCount = mediaCount
		if collectionMetadata.Title != "" {
			s.media.Update(collection.ID, &collectionMetadata)
			collectionUpdated = true
		}
	}

	// Get medias metadata
	for _, media := range *medias {
		s.ScanMedia(media)
	}

	// Scrap collection metadata
	if !collectionUpdated {
		collectionMetadata := scrapers.Scrap(collection.FileName)
		collectionMetadata.MediaCount = mediaCount
		s.media.Update(collection.ID, &collectionMetadata)
	}
	return nil
}

// ScanMedia scans metadata for the media
func (s *MetadataService) ScanMedia(media db.Media) error {
	comic, err := comicfile.New(media.Path)
	if err != nil {
		s.media.Delete(media.ID)
		return nil
	}
	mediaMetadata, err := comic.ExtractMetadata()
	if err != nil {
		return err
	}
	s.media.Update(media.ID, &mediaMetadata)
	return nil
}

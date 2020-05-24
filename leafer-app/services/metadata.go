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
	s.ScanCollectionsLibrary(id)
	s.ScanMediasLibrary(id)
	return nil
}

// ScanMedia scans metadata for the media
func (s *MetadataService) ScanMedia(media db.Media) error {
	log.Printf("Metadata media [%v]", media.Path)

	comic, err := comicfile.New(media.Path)
	if err != nil {
		return err
	}
	mediaMetadata, err := comic.ExtractMetadata()
	if err != nil {
		return err
	}
	s.media.Update(media.ID, &mediaMetadata)

	// Scrap metadata for standalone media
	if media.ParentMediaID == 0 {
		mediaMetadata := scrapers.Scrap(media.FileName)
		s.media.Update(media.ID, &mediaMetadata)
	}
	return nil
}

// ScanMediasLibrary scans metadata for all medias of the library
func (s *MetadataService) ScanMediasLibrary(libraryID uint) error {
	medias, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(libraryID), MediaType: "MEDIA"})
	if err != nil {
		return err
	}
	for _, media := range *medias {
		s.ScanMedia(media)
	}
	return nil
}

// ScanCollection scans metadata for the collection
func (s *MetadataService) ScanCollection(collection db.Media) error {
	log.Printf("Metadata collection [%v]", collection.Path)

	// Count medias into the collection
	mediaCount := s.media.CountSearch(db.SearchMediaInputs{LibraryID: fmt.Sprint(collection.LibraryID), ParentMediaID: fmt.Sprint(collection.ID)})

	// Get collection metadata from the first media metadata
	firstMedia, _ := s.media.GetFirstMediaCollection(collection.LibraryID, collection.ID)
	collectionUpdated := false
	if firstMedia != nil {
		comic, _ := comicfile.New(firstMedia.Path)
		collectionMetadata, _ := comic.ExtractMetadata()
		if collectionMetadata.Title != "" {
			cover, err := comic.ExtractCover(collection.ID)
			if err == nil {
				collectionMetadata.CoverImageLocal = cover
			}
			collectionMetadata.MediaCount = mediaCount
			s.media.Update(collection.ID, &collectionMetadata)
			collectionUpdated = true
		}
	}

	// Scrap collection metadata
	if !collectionUpdated {
		collectionMetadata := scrapers.Scrap(collection.FileName)
		collectionMetadata.MediaCount = mediaCount
		s.media.Update(collection.ID, &collectionMetadata)
	}
	return nil
}

// ScanCollectionsLibrary scans metadata for all collections of the library
func (s *MetadataService) ScanCollectionsLibrary(libraryID uint) error {
	collections, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(libraryID), MediaType: "COLLECTION"})
	if err != nil {
		return err
	}
	for _, collection := range *collections {
		s.ScanCollection(collection)
	}
	return nil
}

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
	s.ScanSeriesLibrary(id)
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
	if media.SerieID == 0 {
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

// ScanSerie scans metadata for the serie
func (s *MetadataService) ScanSerie(serie db.Media) error {
	log.Printf("Metadata serie [%v]", serie.Path)

	// Count medias into the serie
	mediaCount := s.media.CountSearch(db.SearchMediaInputs{LibraryID: fmt.Sprint(serie.LibraryID), SerieID: fmt.Sprint(serie.ID)})

	// Get serie metadata from the first media metadata
	firstMedia, _ := s.media.GetFirstMediaSerie(serie.LibraryID, serie.ID)
	serieUpdated := false
	if firstMedia != nil {
		comic, _ := comicfile.New(firstMedia.Path)
		serieMetadata, _ := comic.ExtractMetadata()
		if serieMetadata.Title != "" {
			cover, err := comic.ExtractCover(serie.ID)
			if err == nil {
				serieMetadata.CoverImageLocal = cover
			}
			serieMetadata.MediaCount = mediaCount
			s.media.Update(serie.ID, &serieMetadata)
			serieUpdated = true
		}
	}

	// Scrap serie metadata
	if !serieUpdated {
		serieMetadata := scrapers.Scrap(serie.FileName)
		serieMetadata.MediaCount = mediaCount
		s.media.Update(serie.ID, &serieMetadata)
	}
	return nil
}

// ScanSeriesLibrary scans metadata for all series of the library
func (s *MetadataService) ScanSeriesLibrary(libraryID uint) error {
	series, err := s.media.Search(db.SearchMediaInputs{LibraryID: fmt.Sprint(libraryID), MediaType: "SERIE"})
	if err != nil {
		return err
	}
	for _, serie := range *series {
		s.ScanSerie(serie)
	}
	return nil
}

package services

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/comicfile"
	"github.com/bpetetot/leafer/services/utils"
	"github.com/jinzhu/gorm"
)

// ScannerService exposes service to scan a library
type ScannerService struct {
	library db.LibraryStore
	media   db.MediaStore
}

// NewScannerService creates a library scanner service instance
func NewScannerService(DB *gorm.DB) ScannerService {
	return ScannerService{
		library: db.NewLibraryStore(DB),
		media:   db.NewMediaStore(DB),
	}
}

// ScanLibrary scans the given library id
func (s *ScannerService) ScanLibrary(id uint) error {
	library, err := s.library.Get(id)
	if err != nil {
		return err
	}

	log.Printf("Scan files for library '%s' [%s]", library.Name, library.Path)

	// Set all current media of the library to "scanning" status
	err = s.media.UpdateMediasLibraryScanningStatus(library.ID, "scanning")
	if err != nil {
		return err
	}

	// Scan the library directory
	s.scanDirectory(library.Path, library, nil)

	// Delete all remaining "scanning" status from the library
	err = s.media.DeleteMediasWithScanningStatus(library.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScannerService) scanDirectory(path string, library *db.Library, parentMedia *db.Media) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	utils.SortFiles(files)
	mediaIndex := 1
	for _, file := range files {
		filename := file.Name()
		if utils.IsHidden(filename) {
			continue
		}

		newPath := filepath.Join(path, filename)
		if file.IsDir() {
			serie, _ := s.createSerie(newPath, file, library)
			if serie != nil {
				log.Printf("Scanning [%s]", newPath)
				s.scanDirectory(newPath, library, serie)
			}
		} else {
			ext := filepath.Ext(filename)
			if utils.Contains(comicfile.ArchiveExt, ext) {
				_, err = s.createMedia(newPath, file, library, parentMedia, mediaIndex)
				if err == nil {
					mediaIndex++
				}
			}
		}
	}
}

func (s *ScannerService) createSerie(path string, info os.FileInfo, library *db.Library) (*db.Media, error) {
	savedMedia := &db.Media{
		Type:           "SERIE",
		Library:        library,
		Path:           path,
		FileName:       info.Name(),
		ScanningStatus: "scanned",
	}

	currentMedia, _ := s.media.ExistMediaPath(library.ID, path)
	if currentMedia != nil {
		log.Printf("Updated serie [%s]", path)
		s.media.Update(currentMedia.ID, savedMedia)
		return s.media.Get(currentMedia.ID)
	}
	log.Printf("Created serie [%s]", path)
	return s.media.Create(savedMedia)
}

func (s *ScannerService) createMedia(path string, info os.FileInfo, library *db.Library, serie *db.Media, mediaIndex int) (*db.Media, error) {
	savedMedia := &db.Media{
		Type:           "MEDIA",
		Library:        library,
		Serie:          serie,
		Path:           path,
		FileName:       info.Name(),
		MediaIndex:     mediaIndex,
		ScanningStatus: "scanned",
	}

	currentMedia, _ := s.media.ExistMediaPath(library.ID, path)
	if currentMedia != nil {
		log.Printf("Updated media [%s]", path)
		s.media.Update(currentMedia.ID, savedMedia)
		return s.media.Get(currentMedia.ID)
	}
	log.Printf("Created media [%s]", path)
	return s.media.Create(savedMedia)
}

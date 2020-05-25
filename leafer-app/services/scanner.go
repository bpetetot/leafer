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
	err = s.media.DeleteMediasLibrary(library.ID)
	if err != nil {
		return err
	}

	s.scanDirectory(library.Path, library, nil)
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
	return s.media.Create(&db.Media{
		Type:     "SERIE",
		Library:  library,
		Path:     path,
		FileName: info.Name(),
	})
}

func (s *ScannerService) createMedia(path string, info os.FileInfo, library *db.Library, serie *db.Media, mediaIndex int) (*db.Media, error) {
	return s.media.Create(&db.Media{
		Type:       "MEDIA",
		Library:    library,
		Serie:      serie,
		Path:       path,
		FileName:   info.Name(),
		MediaIndex: mediaIndex,
	})
}

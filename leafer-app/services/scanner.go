package services

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/utils"
	"github.com/bpetetot/leafer/services/zip"
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
			collection, _ := s.createCollection(newPath, file, library)
			if collection != nil {
				log.Printf("Scanning [%s]", newPath)
				s.scanDirectory(newPath, library, collection)
			}
		} else {
			ext := filepath.Ext(filename)
			if utils.Contains(zip.ArchiveExt, ext) {
				_, err = s.createMedia(newPath, file, library, parentMedia, mediaIndex)
				if err == nil {
					mediaIndex++
				}
			}
		}
	}
}

func (s *ScannerService) createCollection(path string, info os.FileInfo, library *db.Library) (*db.Media, error) {
	return s.media.Create(&db.Media{
		Type:          "COLLECTION",
		Library:       library,
		Path:          path,
		EstimatedName: info.Name(),
	})
}

func (s *ScannerService) createMedia(path string, info os.FileInfo, library *db.Library, collection *db.Media, mediaIndex int) (*db.Media, error) {
	// get basic file info
	basename := info.Name()
	extension := filepath.Ext(basename)
	name := strings.TrimSuffix(basename, extension)
	volume, _ := strconv.Atoi(getVolumeNumber(name))

	// get media info from file
	zipFilesList, err := zip.ListImages(path)
	if err != nil {
		return nil, err
	}

	return s.media.Create(&db.Media{
		Type:          "MEDIA",
		Library:       library,
		ParentMedia:   collection,
		Path:          path,
		MediaIndex:    mediaIndex,
		EstimatedName: getVolumeName(name),
		FileName:      basename,
		FileExtension: extension,
		Volume:        volume,
		PageCount:     len(zipFilesList),
	})
}

func getVolumeNumber(filename string) string {
	re := regexp.MustCompile(`\d+`)
	numbers := re.FindAllString(filename, -1)
	if len(numbers) == 0 {
		return "0"
	}
	return numbers[len(numbers)-1]
}

func getVolumeName(filename string) string {
	number := getVolumeNumber(filename)
	return strings.Replace(filename, number, "", -1)
}

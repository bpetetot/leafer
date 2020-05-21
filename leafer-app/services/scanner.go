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
	s.media.DeleteMediasLibrary(library.ID)
	s.scanDirectory(library.Path, library, nil, 0)
	return nil
}

func (s *ScannerService) scanDirectory(path string, library *db.Library, parentMedia *db.Media, mediaIndex int) int {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return 0
	}

	utils.SortFiles(files)
	for _, file := range files {
		if file.Name()[0:1] == "." {
			continue
		}

		newPath := filepath.Join(path, file.Name())
		if file.IsDir() {
			collection := s.createCollection(newPath, file, library)
			if collection != nil {
				log.Printf("Scanning [%s]", newPath)
				s.scanDirectory(newPath, library, collection, 0)
			}
		} else {
			mediaIndex++
			s.createMedia(newPath, file, library, parentMedia, mediaIndex)
		}
	}
	return mediaIndex
}

func (s *ScannerService) createCollection(path string, info os.FileInfo, library *db.Library) *db.Media {
	collection := db.Media{
		Type:          "COLLECTION",
		Library:       library,
		Path:          path,
		EstimatedName: info.Name(),
	}
	s.media.Create(&collection)
	return &collection
}

func (s *ScannerService) createMedia(path string, info os.FileInfo, library *db.Library, collection *db.Media, mediaIndex int) {
	matched, err := filepath.Match("*.zip", info.Name())
	if !matched || err != nil {
		return
	}

	// get basic file info
	basename := info.Name()
	extension := filepath.Ext(basename)
	name := strings.TrimSuffix(basename, extension)
	volume, _ := strconv.Atoi(getVolumeNumber(name))

	// get media info from file
	zipFilesList, _ := zip.ListImages(path)

	s.media.Create(&db.Media{
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
	return numbers[len(numbers)-1]
}

func getVolumeName(filename string) string {
	number := getVolumeNumber(filename)
	return strings.Replace(filename, number, "", -1)
}

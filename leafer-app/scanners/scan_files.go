package scanners

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/utils"
	"github.com/jinzhu/gorm"
)

// ScanLibrary scans all files from a library
func ScanLibrary(library *db.Library, conn *gorm.DB) {
	log.Printf("Scan files for library '%s' [%s]", library.Name, library.Path)
	db.DeleteLibraryContent(library, conn)
	scanDirectory(library.Path, library, nil, 0, conn)
}

func scanDirectory(path string, library *db.Library, parentMedia *db.Media, mediaIndex int, conn *gorm.DB) int {
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
			collection := createCollection(newPath, file, library, conn)

			if collection != nil {
				mediaCount := scanDirectory(newPath, library, collection, 0, conn)

				conn.Model(&collection).Update(db.Media{MediaCount: mediaCount})

				log.Printf("[%s] %v media", newPath, mediaCount)
			}
		} else {
			mediaIndex++
			createMedia(newPath, file, library, parentMedia, mediaIndex, conn)
		}
	}
	return mediaIndex
}

func createCollection(path string, info os.FileInfo, library *db.Library, conn *gorm.DB) *db.Media {
	collection := &db.Media{
		Type:          "COLLECTION",
		Library:       library,
		Path:          path,
		EstimatedName: info.Name(),
	}
	conn.Create(&collection)
	return collection
}

func createMedia(path string, info os.FileInfo, library *db.Library, collection *db.Media, mediaIndex int, conn *gorm.DB) {
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
	zipFilesList, _ := utils.ListImagesInZip(path)

	media := &db.Media{
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
	}
	conn.Create(&media)
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

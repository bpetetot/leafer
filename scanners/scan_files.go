package scanners

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/bpetetot/leafer/db"
	"github.com/jinzhu/gorm"
)

// ScanFiles scans all files from a library
func ScanFiles(library *db.Library, conn *gorm.DB) error {
	root, err := os.Stat(library.Path)
	if err != nil {
		return errors.New("the library path does not exist anymore")
	}

	db.DeleteLibraryContent(library, conn)

	scanContent(library.Path, root, library, nil, conn)

	return nil
}

func scanContent(path string, info os.FileInfo, library *db.Library, parentMedia *db.Media, conn *gorm.DB) {
	if info.IsDir() {
		files, _ := ioutil.ReadDir(path)
		for _, file := range files {
			filename := file.Name()
			newPath := filepath.Join(path, filename)

			if filename[0:1] == "." {
				continue
			}

			var curMedia *db.Media = parentMedia
			if parentMedia == nil {
				curMedia = &db.Media{Library: library, EstimatedName: file.Name()}
				conn.Create(&curMedia)
			}

			scanContent(newPath, file, library, curMedia, conn)
		}
	} else {
		media, err := createMediaFromFile(path, info)
		if err != nil {
			return
		}

		media.Library = library
		media.ParentMedia = parentMedia
		conn.Create(&media)
	}
}

func createMediaFromFile(path string, info os.FileInfo) (*db.Media, error) {
	matched, err := filepath.Match("*.zip", info.Name())
	if !matched || err != nil {
		return nil, errors.New("does not match to analyzed extensions")
	}

	basename := info.Name()
	extension := filepath.Ext(basename)
	name := strings.TrimSuffix(basename, extension)
	volume, _ := strconv.Atoi(getVolumeNumber(name))

	return &db.Media{
		EstimatedName: getVolumeName(name),
		Volume:        volume,
		FilePath:      filepath.Join(path, basename),
		FileDir:       path,
		FileName:      basename,
		FileExtension: extension,
	}, nil
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

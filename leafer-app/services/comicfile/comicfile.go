package comicfile

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/mholt/archiver/v3"

	"github.com/bpetetot/leafer/services/utils"
)

var zipArchiver = archiver.Zip{ContinueOnError: true}
var rarArchiver = archiver.Rar{ContinueOnError: true}

var zipExt = []string{".zip", ".cbz"}
var rarExt = []string{".rar", ".cbr"}
var imageExt = []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"}
var metadataFiles = []string{"ComicInfo.xml"}

// ArchiveExt contains handled archive extensions
var ArchiveExt = append(zipExt, rarExt...)

type fileArchiver interface {
	Walk(archive string, walkFn archiver.WalkFunc) error
}

func getFileArchiver(src string) (fileArchiver, error) {
	ext := filepath.Ext(src)
	if utils.Contains(zipExt, ext) {
		return &zipArchiver, nil
	}
	if utils.Contains(rarExt, ext) {
		return &rarArchiver, nil
	}
	return nil, errors.New("No archiver found for file extension")
}

// ComicFile is a comic file archive
type ComicFile struct {
	filepath string
	archiver fileArchiver
}

// New creates a readable ComicFile
func New(filepath string) (*ComicFile, error) {
	fileArchiver, err := getFileArchiver(filepath)
	if err != nil {
		return nil, err
	}
	return &ComicFile{
		filepath: filepath,
		archiver: fileArchiver,
	}, nil
}

// ListImages lists ordered images within a comic file
func (c *ComicFile) ListImages() ([]string, error) {
	var filenames []string
	err := c.archiver.Walk(c.filepath, func(f archiver.File) error {
		if isImageFile(f) {
			filenames = append(filenames, f.Name())
		}
		return nil
	})

	sort.Sort(utils.Natural(filenames))

	return filenames, err
}

// ExtractImage extracts a specific image index in the comic file into the given io.Writer
func (c *ComicFile) ExtractImage(index int, w io.Writer) error {
	filenames, err := c.ListImages()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(filenames) {
		return errors.New("index not found")
	}

	err = c.archiver.Walk(c.filepath, func(f archiver.File) error {
		if f.Name() == filenames[index] {
			_, err = io.Copy(w, f)
			return err
		}
		return nil
	})

	return err
}

// ComicMetadata comic metadata extracted from the file
type ComicMetadata struct {
	PageCount int
	Volume    int
}

// ExtractMetadata extracts metadata of the comic file into the given io.Writer
func (c *ComicFile) ExtractMetadata() (ComicMetadata, error) {
	basename := filepath.Base(c.filepath)
	volume := getVolumeNumber(basename)
	pageCount := 0
	err := c.archiver.Walk(c.filepath, func(f archiver.File) error {
		if isImageFile(f) {
			pageCount++
		}
		return nil
	})

	return ComicMetadata{PageCount: pageCount, Volume: volume}, err
}

func isImageFile(file os.FileInfo) bool {
	ext := filepath.Ext(file.Name())
	return utils.Contains(imageExt, ext) && !utils.IsHidden(file.Name()) && !file.IsDir()
}

func isMetadataFile(file os.FileInfo) bool {
	return utils.Contains(metadataFiles, file.Name())
}

func getVolumeNumber(filename string) int {
	re := regexp.MustCompile(`\b(\d{1,3})\b`)
	numbers := re.FindAllString(filename, -1)
	if len(numbers) == 0 {
		return 0
	}
	number, err := strconv.Atoi(numbers[len(numbers)-1])
	if err != nil {
		return 0
	}
	return number
}

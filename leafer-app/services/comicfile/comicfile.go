package comicfile

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/mholt/archiver/v3"
	"github.com/nfnt/resize"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/comicfile/metadata"
	"github.com/bpetetot/leafer/services/utils"
)

var zipExt = []string{".zip", ".cbz"}
var rarExt = []string{".rar", ".cbr"}
var imageExt = []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"}

// ArchiveExt contains handled archive extensions
var ArchiveExt = append(zipExt, rarExt...)

type fileArchiver interface {
	Walk(archive string, walkFn archiver.WalkFunc) error
}

func getFileArchiver(src string) (fileArchiver, error) {
	ext := filepath.Ext(src)
	if utils.Contains(zipExt, ext) {
		return &archiver.Zip{ContinueOnError: true}, nil
	}
	if utils.Contains(rarExt, ext) {
		return &archiver.Rar{ContinueOnError: true}, nil
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

// ExtractMetadata extracts metadata of the comic file
func (c *ComicFile) ExtractMetadata() (db.Media, error) {
	basename := filepath.Base(c.filepath)
	volume := getVolumeNumber(basename)
	pageCount := 0

	var media = &db.Media{}
	err := c.archiver.Walk(c.filepath, func(f archiver.File) error {
		if isImageFile(f) {
			pageCount++
		} else if metadata.IsMetadataFile(f) {
			media, _ = metadata.ReadMetadata(f.Name(), f)
		}
		return nil
	})

	if media.PageCount == 0 {
		media.PageCount = pageCount
	}

	if media.Volume == 0 {
		media.Volume = volume
	}

	return *media, err
}

// ExtractCover extracts and resize cover image to be saved
func (c *ComicFile) ExtractCover(id uint) (string, error) {
	folder := filepath.Join(".metadata", fmt.Sprint("serie-", id))
	os.MkdirAll(folder, os.ModePerm)

	buf := new(bytes.Buffer)
	err := c.ExtractImage(0, buf)
	if err != nil {
		return "", err
	}
	originalImage, _, err := image.Decode(buf)
	if err != nil {
		return "", err
	}
	resizedImage := resize.Resize(224, 0, originalImage, resize.Lanczos3)

	coverPath := filepath.Join(folder, "cover.jpg")
	f, err := os.Create(coverPath)
	defer f.Close()
	if err != nil {
		return "", err
	}

	err = jpeg.Encode(f, resizedImage, nil)
	return fmt.Sprint("/metadata/serie-", id, "/cover.jpg"), err
}

func isImageFile(file os.FileInfo) bool {
	ext := filepath.Ext(file.Name())
	return utils.Contains(imageExt, ext) && !utils.IsHidden(file.Name()) && !file.IsDir()
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

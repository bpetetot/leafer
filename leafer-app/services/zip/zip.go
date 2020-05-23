package zip

import (
	"errors"
	"io"
	"path/filepath"
	"sort"

	"github.com/mholt/archiver/v3"

	"github.com/bpetetot/leafer/services/utils"
)

var zipArchiver = archiver.Zip{ContinueOnError: true}
var rarArchiver = archiver.Rar{ContinueOnError: true}

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
		return &zipArchiver, nil
	}
	if utils.Contains(rarExt, ext) {
		return &rarArchiver, nil
	}
	return nil, errors.New("No archiver found for file extension")
}

// ListImages will list all images within the zip archive
func ListImages(src string) ([]string, error) {
	var filenames []string
	fileArchiver, err := getFileArchiver(src)
	if err != nil {
		return filenames, err
	}

	err = fileArchiver.Walk(src, func(f archiver.File) error {
		name := f.Name()
		ext := filepath.Ext(name)
		matched := utils.Contains(imageExt, ext) && !utils.IsHidden(name) && !f.IsDir()
		if matched {
			filenames = append(filenames, name)
		}
		return nil
	})

	sort.Sort(utils.Natural(filenames))

	return filenames, err
}

func extractFile(src string, filename string, w io.Writer) error {
	fileArchiver, err := getFileArchiver(src)
	if err != nil {
		return err
	}

	err = fileArchiver.Walk(src, func(f archiver.File) error {
		if f.Name() == filename {
			_, err = io.Copy(w, f)
			return err
		}
		return nil
	})

	return nil
}

// ExtractImage extracts a specific image index in the archive
// and stream it to the given io.Writer
func ExtractImage(src string, index int, w io.Writer) error {
	filenames, err := ListImages(src)
	if err != nil {
		return err
	}

	if index < 0 || index >= len(filenames) {
		return errors.New("index not found")
	}

	return extractFile(src, filenames[index], w)
}

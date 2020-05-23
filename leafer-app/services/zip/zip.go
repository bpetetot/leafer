package zip

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
	"sort"

	"github.com/bpetetot/leafer/services/utils"
)

// ListImages will list all images within the zip archive
func ListImages(src string) ([]string, error) {
	var extensions = []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"}
	var filenames []string

	reader, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		info := file.FileInfo()
		ext := filepath.Ext(info.Name())
		matched := utils.Contains(extensions, ext) && !utils.IsHidden(info.Name()) && !info.IsDir()

		if !matched {
			continue
		}
		filenames = append(filenames, file.Name)
	}

	sort.Sort(utils.Natural(filenames))

	return filenames, nil
}

// StreamImage Unzip a specific image in the archive and stream it to the
// given io.Writer
func StreamImage(src string, index int, w io.Writer) error {
	filenames, err := ListImages(src)
	if err != nil {
		return err
	}

	if index < 0 || index >= len(filenames) {
		return errors.New("index not found")
	}

	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		if file.Name == filenames[index] {
			reader, err := file.Open()
			if err != nil {
				return err
			}
			defer reader.Close()

			_, err = io.Copy(w, reader)
			return err
		}
	}

	return nil
}

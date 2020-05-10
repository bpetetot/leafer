package utils

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
	"sort"
)

// ListImagesInZip will list all images within the zip archive
func ListImagesInZip(src string) ([]string, error) {
	var extensions = []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"}
	var filenames []string

	reader, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer reader.Close()

	for _, file := range reader.File {
		ext := filepath.Ext(file.Name)
		matched := Contains(extensions, ext)

		if !matched {
			continue
		}
		filenames = append(filenames, file.Name)
	}

	sort.Sort(Natural(filenames))

	return filenames, nil
}

// StreamImageFromZip Unzip a specific image in the archive and stream it to the
// given io.Writer
func StreamImageFromZip(src string, index int, w io.Writer) error {
	filenames, err := ListImagesInZip(src)
	if err != nil {
		return err
	}
	if index >= len(filenames) || index < 0 {
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

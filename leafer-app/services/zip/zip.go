package zip

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"

	"github.com/mholt/archiver/v3"

	"github.com/bpetetot/leafer/services/utils"
)

var z = archiver.Zip{
	ContinueOnError: true,
}

// ListImages will list all images within the zip archive
func ListImages(src string) ([]string, error) {
	var extensions = []string{".jpg", ".jpeg", ".png", ".bmp", ".gif"}
	var filenames []string

	err := z.Walk(src, func(f archiver.File) error {
		zfh, ok := f.Header.(zip.FileHeader)
		if ok {
			info := zfh.FileInfo()
			ext := filepath.Ext(info.Name())
			matched := utils.Contains(extensions, ext) && !utils.IsHidden(info.Name()) && !info.IsDir()

			if matched {
				filenames = append(filenames, zfh.Name)
			}
		}
		return nil
	})

	return filenames, err
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

	err = z.Walk(src, func(f archiver.File) error {
		zfh, ok := f.Header.(zip.FileHeader)
		if ok && zfh.Name == filenames[index] {
			_, err = io.Copy(w, f)
			return err
		}
		return nil
	})

	return nil
}

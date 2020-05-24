package metadata

import (
	"errors"
	"io"
	"os"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/services/utils"
)

var metadataFiles = []string{"ComicInfo.xml"}

// IsMetadataFile returns true if file corresponds to a handled metadata files
func IsMetadataFile(file os.FileInfo) bool {
	return utils.Contains(metadataFiles, file.Name())
}

// ReadMetadata reads the file content
func ReadMetadata(filename string, reader io.Reader) (*db.Media, error) {
	if filename == "ComicInfo.xml" {
		return readComicInfo(reader)
	}
	return nil, errors.New("Metadata file not found")
}

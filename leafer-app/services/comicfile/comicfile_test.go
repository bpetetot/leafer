package comicfile

import (
	"bytes"
	"testing"

	"github.com/mholt/archiver/v3"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return the ZIP archiver for zip extensions", func(t *testing.T) {
		comicFile, _ := New("file.zip")
		assert.IsType(&archiver.Zip{}, comicFile.archiver)
		comicFile, _ = New("file.cbz")
		assert.IsType(&archiver.Zip{}, comicFile.archiver)
	})

	t.Run("should return the RAR archiver for rar extensions", func(t *testing.T) {
		comicFile, _ := New("file.rar")
		assert.IsType(&archiver.Rar{}, comicFile.archiver)
		comicFile, _ = New("file.cbr")
		assert.IsType(&archiver.Rar{}, comicFile.archiver)
	})

	t.Run("should return an error if no archiver found for unknown extension", func(t *testing.T) {
		_, err := New("file.bob")
		assert.Error(err)
	})
}

func Test_ListImages(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return an error if archive doesn't exist", func(t *testing.T) {
		comicFile, _ := New("testdata/unknown.zip")
		_, err := comicFile.ListImages()
		assert.Error(err)
	})

	t.Run("should list 5 ordered images in the archive", func(t *testing.T) {
		comicFile, _ := New("testdata/archive.zip")
		images, _ := comicFile.ListImages()
		expected := []string{"test1.jpg", "test2.jpeg", "test3.png", "test4.bmp", "test5.gif"}
		assert.Exactly(expected, images)
	})
}

func Test_ExtractImage(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return an error if archive doesn't exist", func(t *testing.T) {
		var b bytes.Buffer
		comicFile, _ := New("testdata/unknown.zip")
		err := comicFile.ExtractImage(0, &b)
		assert.Error(err)
	})

	t.Run("should return an error if index doesn't exist", func(t *testing.T) {
		var b bytes.Buffer
		comicFile, _ := New("testdata/archive.zip")
		err := comicFile.ExtractImage(-1, &b)
		assert.Error(err)
		err = comicFile.ExtractImage(6, &b)
		assert.Error(err)
	})

	t.Run("should copy the image of the given index in the writer", func(t *testing.T) {
		var b bytes.Buffer
		comicFile, _ := New("testdata/archive.zip")
		err := comicFile.ExtractImage(1, &b)
		assert.NoError(err)
		actual := b.String()
		assert.Equal("test2", actual)
	})
}

func Test_ExtractMetadata(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return minimal file metadata", func(t *testing.T) {
		comicFile, _ := New("testdata/archive.zip")
		metadata, _ := comicFile.ExtractMetadata()
		assert.Equal(ComicMetadata{PageCount: 5, Volume: 0}, metadata)
	})
}

func Test_getVolumeNumber(t *testing.T) {
	tests := []struct {
		filename string
		expected int
	}{
		// success
		{filename: "1.jpg", expected: 1},
		{filename: "01.jpg", expected: 1},
		{filename: "001.jpg", expected: 1},
		{filename: "name-002.jpg", expected: 2},
		{filename: "name 003.jpg", expected: 3},
		{filename: "name 004.jpg", expected: 4},
		{filename: "name 005 (2020).jpg", expected: 5},
		{filename: "name 999.jpg", expected: 999},
		{filename: "name 10 999.jpg", expected: 999},
		// fails
		{filename: "name.jpg", expected: 0},
		{filename: "name_004.jpg", expected: 0},
		{filename: "name 1004.jpg", expected: 0},
	}
	for _, test := range tests {
		actual := getVolumeNumber(test.filename)
		assert.Equal(t, test.expected, actual)
	}
}

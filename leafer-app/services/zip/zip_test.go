package zip

import (
	"bytes"
	"testing"

	"github.com/mholt/archiver/v3"
	"github.com/stretchr/testify/assert"
)

func Test_getFileArchiver(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return the ZIP archiver for zip extensions", func(t *testing.T) {
		fileArchiver, _ := getFileArchiver("file.zip")
		assert.IsType(&archiver.Zip{}, fileArchiver)
		fileArchiver, _ = getFileArchiver("file.cbz")
		assert.IsType(&archiver.Zip{}, fileArchiver)
	})

	t.Run("should return the RAR archiver for rar extensions", func(t *testing.T) {
		fileArchiver, _ := getFileArchiver("file.rar")
		assert.IsType(&archiver.Rar{}, fileArchiver)
		fileArchiver, _ = getFileArchiver("file.cbr")
		assert.IsType(&archiver.Rar{}, fileArchiver)
	})

	t.Run("should return an error if no archiver found for unknown extension", func(t *testing.T) {
		_, err := getFileArchiver("file.bob")
		assert.Error(err)
	})
}

func Test_ListImages(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return an error if archive doesn't exist", func(t *testing.T) {
		_, err := ListImages("testdata/unknown.zip")
		assert.Error(err)
	})

	t.Run("should list 5 ordered images in the archive", func(t *testing.T) {
		images, _ := ListImages("testdata/archive.zip")
		expected := []string{"test1.jpg", "test2.jpeg", "test3.png", "test4.bmp", "test5.gif"}
		assert.Exactly(expected, images)
	})
}

func Test_ExtractImage(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return an error if archive doesn't exist", func(t *testing.T) {
		var b bytes.Buffer
		err := ExtractImage("testdata/unknown.zip", 0, &b)
		assert.Error(err)
	})

	t.Run("should return an error if index doesn't exist", func(t *testing.T) {
		var b bytes.Buffer
		err := ExtractImage("testdata/archive.zip", -1, &b)
		assert.Error(err)
		err = ExtractImage("testdata/archive.zip", 6, &b)
		assert.Error(err)
	})

	t.Run("should copy the image of the given index in the writer", func(t *testing.T) {
		var b bytes.Buffer
		err := ExtractImage("testdata/archive.zip", 1, &b)
		assert.NoError(err)
		actual := b.String()
		assert.Equal("test2", actual)
	})
}

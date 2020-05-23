package zip

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListImages(t *testing.T) {
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

func TestStreamImage(t *testing.T) {
	assert := assert.New(t)

	t.Run("should return an error if archive doesn't exist", func(t *testing.T) {
		var b bytes.Buffer
		err := StreamImage("testdata/unknown.zip", 0, &b)
		assert.Error(err)
	})

	t.Run("should return an error if index doesn't exist", func(t *testing.T) {
		var b bytes.Buffer
		err := StreamImage("testdata/archive.zip", -1, &b)
		assert.Error(err)
		err = StreamImage("testdata/archive.zip", 6, &b)
		assert.Error(err)
	})

	t.Run("should copy the image of the given index in the writer", func(t *testing.T) {
		var b bytes.Buffer
		err := StreamImage("testdata/archive.zip", 1, &b)
		assert.NoError(err)
		actual := b.String()
		assert.Equal("test2", actual)
	})
}

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsHidden(t *testing.T) {
	result := IsHidden("file.txt")
	assert.Equal(t, false, result, "should not be a hidden file")

	result = IsHidden(".file.txt")
	assert.Equal(t, true, result, "should be a hidden file")
}

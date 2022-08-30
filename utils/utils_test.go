package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFilesAndDirs(t *testing.T) {
	fs, ds, err := GetFilesAndDirs("../")
	assert.Equal(t, err, nil)
	fmt.Println(ds, fs)
}

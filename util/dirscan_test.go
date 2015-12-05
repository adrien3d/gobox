package util

import (
	"testing"
)

func TestDirscan(t *testing.T) {
	if _, err := scanDir("./"); err != nil {
		t.Error(err)
	}
}

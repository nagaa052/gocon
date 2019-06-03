package search_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nagaa052/gocon/pkg/search"
)

func TestWalk(t *testing.T) {
	errMessage := `
	The quantity found is different.
		expected: %d
		actual: %d
	`
	cases := []struct {
		ext      []string
		expected int
	}{
		{ext: []string{".jpg"}, expected: 4},
		{ext: []string{".png"}, expected: 2},
		{ext: []string{".jpg", ".png"}, expected: 6},
	}

	path := filepath.Join("../testdata")
	for _, c := range cases {
		actual := 0
		search.WalkWithExtHandle(path, c.ext, func(path string, info os.FileInfo, err error) {
			actual++
		})
		if c.expected != actual {
			t.Error(fmt.Sprintf(errMessage, c.expected, actual))
		}
	}
}

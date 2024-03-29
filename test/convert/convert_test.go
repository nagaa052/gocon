package convert_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nagaa052/gocon/pkg/convert"
)

func TestToJpeg(t *testing.T) {
	errMessage := `
	Expected and actual are different.
		expected: %v
		actual: %v
	`

	srcBase, err := filepath.Abs("../testdata/dirA")
	if err != nil {
		t.Error(err)
	}

	destDir := filepath.Join(srcBase, "out")

	cases := []struct {
		srcPath  string
		expected string
		isError  bool
	}{
		{
			srcPath:  filepath.Join(srcBase, "1.png"),
			expected: filepath.Join(destDir, "1.jpeg"),
			isError:  false,
		},
		{
			srcPath:  filepath.Join(srcBase, "1.png"),
			expected: "",
			isError:  true,
		},
	}

	for _, c := range cases {
		con, err := convert.New(c.srcPath, destDir)
		if err != nil {
			t.Error(err)
		}

		actual, err := con.ToJpeg(&convert.JpegOptions{})
		if c.isError && err == nil {
			t.Error(err)
		}

		if actual != c.expected {
			t.Error(fmt.Sprintf(errMessage, c.expected, actual))
		}
	}

	if _, err := os.Stat(destDir); err == nil {
		if err := os.RemoveAll(destDir); err != nil {
			fmt.Println(err)
		}
	}
}

func TestToPng(t *testing.T) {
	errMessage := `
	Expected and actual are different.
		expected: %v
		actual: %v
	`

	sourceBase, err := filepath.Abs("../testdata")
	if err != nil {
		t.Error(err)
	}

	outDir := filepath.Join(sourceBase, "out")

	cases := []struct {
		sourcePath string
		expected   string
		isError    bool
	}{
		{
			sourcePath: filepath.Join(sourceBase, "1.jpg"),
			expected:   filepath.Join(outDir, "1.png"),
			isError:    false,
		},
		{
			sourcePath: filepath.Join(sourceBase, "1.jpg"),
			expected:   "",
			isError:    true,
		},
	}
	for _, c := range cases {
		con, err := convert.New(c.sourcePath, outDir)
		if err != nil {
			t.Error(err)
		}

		actual, err := con.ToPng()
		if c.isError && err == nil {
			t.Error(err)
		}

		if actual != c.expected {
			t.Error(fmt.Sprintf(errMessage, c.expected, actual))
		}
	}

	if _, err := os.Stat(outDir); err == nil {
		if err := os.RemoveAll(outDir); err != nil {
			fmt.Println(err)
		}
	}
}

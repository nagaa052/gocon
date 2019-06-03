package gocon_test

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nagaa052/gocon/internal/gocon"
)

func TestGetExtentions(t *testing.T) {
	errMessage := `
	Incorrect extension.
		expected: %s
		actual: %s
	`
	expected := []string{".hogehoge"}

	f := "hogehoge"
	i := gocon.ImgFormat(f)
	actual := i.GetExtentions()
	if len(actual[0]) <= 0 {
		t.Error(fmt.Sprintf(errMessage, expected, actual))
	}

	if expected[0] != actual[0] {
		t.Error(fmt.Sprintf(errMessage, expected, actual))
	}

}

func TestRun(t *testing.T) {
	errMessage := `
	It was not converted properly.
		Error: %v
	`
	streamErrMessage := `
	Convert Error.
		%v
	`

	srcDir, err := filepath.Abs(filepath.Join("../testdata"))
	if err != nil {
		t.Error(err)
	}
	destDir, err := filepath.Abs("out")
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		options  gocon.Options
		expected int
		isError  bool
	}{
		{
			options: gocon.Options{
				FromFormat: gocon.JPEG,
				ToFormat:   gocon.PNG,
				DestDir:    destDir,
			},
			expected: gocon.ExitOK,
			isError:  false,
		},
	}

	for _, c := range cases {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		g, err := gocon.New(srcDir, c.options, outStream, errStream)
		if !c.isError && err != nil {
			t.Error(fmt.Sprintf("gocon make error : %+v", err))
		}

		actual := g.Run()

		if _, err := os.Stat(destDir); err == nil {
			if err := os.RemoveAll(destDir); err != nil {
				fmt.Println(err)
			}
		}

		if c.expected != actual {
			t.Error(fmt.Sprintf(errMessage, c.expected, actual))
		}

		if errStream.String() != "" {
			t.Error(fmt.Sprintf(streamErrMessage, errStream.String()))
		}

		messageFmt := strings.Replace(gocon.SuccessConvertFileMessageFmt, "\n", "", -1)
		expectedMessage := fmt.Sprintf(messageFmt, destDir)
		if !strings.Contains(outStream.String(), expectedMessage) {
			t.Error("Conversion failed")
		}
	}
}

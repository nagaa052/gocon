package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nagaa052/gocon/internal/gocon"
)

var imageFormat string
var destDir string
var outStream io.Writer = os.Stdout
var errStream io.Writer = os.Stderr

func init() {
	flag.StringVar(&imageFormat,
		"f",
		gocon.DefaultOptions.FromFormat.String()+":"+gocon.DefaultOptions.ToFormat.String(),
		"Convert image format. The input format is [In]:[Out]. Default jpeg:png")
	flag.StringVar(&destDir,
		"d",
		gocon.DefaultOptions.DestDir,
		"Destination directory. Default "+gocon.DefaultOptions.DestDir)
}

func main() {

	flag.Usage = usage
	flag.Parse()

	if opt, err := parseOptions(); err != nil {
		fmt.Printf("%s\n", err.Error())
		usage()
	} else {
		for _, srcDir := range flag.Args() {
			gc, err := gocon.New(srcDir, opt, outStream, errStream)
			if err != nil {
				fmt.Printf("%v\n", err.Error())
				os.Exit(1)
			}

			os.Exit(gc.Run())
		}
	}
}

func parseOptions() (gocon.Options, error) {
	format := strings.Split(imageFormat, ":")

	if len(format) != 2 {
		return gocon.Options{}, fmt.Errorf("invalid format")
	}

	from := gocon.ImgFormat(format[0])
	to := gocon.ImgFormat(format[1])

	if !from.Exist() || !to.Exist() {
		return gocon.Options{}, fmt.Errorf("Format that does not exist is an error")
	}

	if from == to {
		return gocon.Options{}, fmt.Errorf("The same format specification is an error")
	}

	return gocon.Options{
		FromFormat: from,
		ToFormat:   to,
		DestDir:    destDir,
	}, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `
gocon is a tool for ...
Usage:
  gocon [option] <directory path>
Options:
`)
	flag.PrintDefaults()
}

package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Mode defines the shapes used when transforming images.
type Mode int

var Mock1 bool
var Mock2 bool
var Mock3 bool
var Mock4 bool
var Mock5 bool
var Mock6 bool

// Mode supported by the primitive package.
const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

// This function is used to select the mode. By Default ModeTriangle will be used.
func WithMode(mode Mode) func() []string {
	return func() []string {
		// fmt.Println([]string{"-m", fmt.Sprintf("%d", mode)})
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

// This function will transform the image using primitive package
func Transform(image io.Reader, ext string, NumShapes int, opts ...func() []string) (io.Reader, error) {

	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}

	in, err := tempFile("in_", ext)
	if err != nil || Mock1 {
		return nil, err
	}
	defer os.Remove(in.Name())

	out, err := tempFile("in_", ext)
	if err != nil || Mock2 {
		return nil, err
	}
	defer os.Remove(out.Name())

	// read input into image file
	_, err = io.Copy(in, image)
	if err != nil || Mock3 {
		return nil, err
	}

	stdCombo, err := Primitive(in.Name(), out.Name(), NumShapes, args...)
	if err != nil || Mock4 {
		return nil, err
	}
	fmt.Println(stdCombo)

	// read out into a reader, return reader, delete out
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil || Mock5 {
		return nil, err
	}
	return b, nil
}

// It will create an image using primitive package with different shapes from an input image
func Primitive(inputFile, outputFile string, numShapes int, args ...string) (string, error) {
	// fmt.Println(args)
	argStr := fmt.Sprintf("-i %s -o %s -n %d", inputFile, outputFile, numShapes)
	args = append(strings.Fields(argStr), args...)
	cmd := exec.Command("primitive", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

// create the temporary file to store images
func tempFile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", prefix)
	if err != nil || Mock6 {
		return nil, err
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

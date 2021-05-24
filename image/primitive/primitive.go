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

// Withmode is an option for the transform function that will define
// the mode you want to use. By Default ModeTriangle will be used.
func WithMode(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

// Transform will the provided image and apply a primitive
// transformation to it, then return a reader to the resulting image.
func Transform(image io.Reader, ext string, NumShapes int, opts ...func() []string) (io.Reader, error) {

	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}

	in, err := tempFile("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())

	out, err := tempFile("in_", ext)
	if err != nil {
		return nil, err
	}
	defer os.Remove(out.Name())

	// read input into image file
	_, err = io.Copy(in, image)
	if err != nil {
		return nil, err
	}

	// Run primitive w/ -i in.Name() -o out.Name()
	stdCombo, err := Primitive(in.Name(), out.Name(), NumShapes, args...)
	if err != nil {
		return nil, err
	}
	fmt.Println(stdCombo)

	// read out into a reader, return reader, delete out
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Primitive(inputFile, outputFile string, numShapes int, args ...string) (string, error) {
	argStr := fmt.Sprintf("-i %s -o %s -n %d", inputFile, outputFile, numShapes)
	args = append(strings.Fields(argStr), args...)
	cmd := exec.Command("primitive", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func tempFile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("", prefix)
	if err != nil {
		return nil, err
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

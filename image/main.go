package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/image/primitive"
)

var temp bool

var Mock1 bool
var Mock2 bool
var Mock3 bool
var Mock4 bool
var Mock5 bool
var Mock6 bool
var Mock7 bool
var Mock8 bool
var Mock9 bool
var Mock10 bool
var Mock11 bool

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", BaseHandler)
	mux.HandleFunc("/upload", UploadHandler)
	mux.HandleFunc("/modify/", ModifyHandler)

	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))

	if temp {
		server := &http.Server{Addr: ":3000", Handler: mux}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	} else {
		log.Fatal(http.ListenAndServe(":3000", mux))
	}

}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	html := `<html><body>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="image">
				<button type="submit">Upload Image</button>
			</form>
			</body></html>`
	fmt.Fprint(w, html)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil || Mock9 {
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)[1:]
	onDisk, err := Tempfile("", ext)
	if err != nil || Mock10 {
		return
	}
	defer onDisk.Close()
	_, err = io.Copy(onDisk, file)
	if err != nil || Mock11 {
		return
	}
	http.Redirect(w, r, "/modify/"+filepath.Base(onDisk.Name()), http.StatusFound)
}

func ModifyHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("./img/" + filepath.Base(r.URL.Path))
	f, err := os.Open("./img/" + filepath.Base(r.URL.Path))
	if err != nil {
		return
	}
	defer f.Close()

	ext := filepath.Ext(f.Name())[1:]

	modeStr := r.FormValue("mode")
	if modeStr == "" {
		renderModeChoices(w, r, f, ext)
		return
	}

	mode, err := strconv.Atoi(modeStr)
	if err != nil {
		return
	}

	nStr := r.FormValue("n")
	if nStr == "" {
		renderNumShapeChoices(w, r, f, ext, primitive.Mode(mode))
		return
	}

	numShapes, err := strconv.Atoi(nStr)
	if err != nil {
		return
	}
	_ = numShapes
	http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
}

type genOpts struct {
	N int
	M primitive.Mode
}

func renderNumShapeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string, mode primitive.Mode) error {
	opts := []genOpts{
		{N: 10, M: mode},
		{N: 20, M: mode},
		{N: 30, M: mode},
		{N: 40, M: mode},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err != nil || Mock1 {
		return err
	}

	html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name      string
		Mode      primitive.Mode
		NumShapes int
	}
	var data []dataStruct
	for i, img := range imgs {
		data = append(data, dataStruct{
			Name:      filepath.Base(img),
			Mode:      opts[i].M,
			NumShapes: opts[i].N,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil || Mock2 {
		return err
	}
	return nil
}

func renderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) error {
	opts := []genOpts{
		{N: 10, M: primitive.ModeCircle},
		{N: 10, M: primitive.ModeBeziers},
		{N: 10, M: primitive.ModePolygon},
		{N: 10, M: primitive.ModeCombo},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err != nil || Mock3 {
		return err
	}

	html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name string
		Mode primitive.Mode
	}
	var data []dataStruct
	for i, img := range imgs {
		data = append(data, dataStruct{
			Name: filepath.Base(img),
			Mode: opts[i].M,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil || Mock4 {
		return err
	}
	return nil
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	for _, opt := range opts {
		rs.Seek(0, 0)
		f, err := genImage(rs, ext, opt.N, opt.M)
		if err != nil || Mock5 {
			return nil, err
		}
		ret = append(ret, f)
	}
	return ret, nil
}

func genImage(r io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {

	out, err := primitive.Transform(r, ext, numShapes, primitive.WithMode(mode))
	if err != nil || Mock6 {
		return "", err
	}

	outFile, err := Tempfile("", ext)
	if err != nil || Mock7 {
		return "", err
	}

	defer outFile.Close()
	io.Copy(outFile, out)
	return outFile.Name(), nil

}

func Tempfile(prefix, ext string) (*os.File, error) {

	in, err := ioutil.TempFile("./img/", prefix)
	if err != nil || Mock8 {
		return nil, err
	}

	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))

}

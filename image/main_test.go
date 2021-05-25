package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/abhishek-devani/Gophercises/go/src/github.com/abhishek-devani/Gophercises/image/primitive"
)

func server() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", BaseHandler)
	mux.HandleFunc("/upload", UploadHandler)
	mux.HandleFunc("/modify/", ModifyHandler)

	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))

	return mux

}

func TestMain(t *testing.T) {
	temp = true
	main()
	temp = false
}

func TestBaseHandler(t *testing.T) {

	statusCode := 200
	testCaseName := "test1"
	req := httptest.NewRequest("GET", "http://localhost:3000/", nil)

	rec := httptest.NewRecorder()

	BaseHandler(rec, req)

	res := rec.Result()
	if res.StatusCode != statusCode {
		t.Errorf("Test case Number: %v Expected %v , Actual status %v", testCaseName, statusCode, res.StatusCode)
	}
}

func TestUploadHandler(t *testing.T) {

	testServer := httptest.NewServer(server())
	defer testServer.Close()

	path := "tmp/input.jpeg"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "http://localhost:3000/upload", body)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res := httptest.NewRecorder()

	UploadHandler(res, req)

	Mock9 = true
	UploadHandler(res, req)
	Mock9 = false

	Mock10 = true
	UploadHandler(res, req)
	Mock10 = false

	Mock11 = true
	UploadHandler(res, req)
	Mock11 = false

	if res.Code != http.StatusFound {
		t.Errorf("Expected %v , Actual status %v", http.StatusFound, res.Code)
	}

}

func TestModifyHandler(t *testing.T) {
	testServer := httptest.NewServer(server())
	defer testServer.Close()

	req := httptest.NewRequest("GET", "http://localhost:3000/068525259.jpeg", nil)
	rec := httptest.NewRecorder()

	ModifyHandler(rec, req)

}

func TestRenderNumShapeChoices(t *testing.T) {

	testServer := httptest.NewServer(server())
	defer testServer.Close()

	req := httptest.NewRequest("GET", "http://localhost:3000/modify/", nil)
	rec := httptest.NewRecorder()

	var i io.ReadSeeker
	i, _ = os.Open("tmp/input.jpeg")

	err := renderNumShapeChoices(rec, req, i, "jpeg", primitive.ModeBeziers)
	if err != nil {
		panic(err)
	}

	Mock1 = true
	err = renderNumShapeChoices(rec, req, i, "jpeg", primitive.ModeBeziers)
	if err != nil {
		panic(err)
	}
	Mock1 = false

	Mock2 = true
	err = renderNumShapeChoices(rec, req, i, "jpeg", primitive.ModeBeziers)
	if err != nil {
		panic(err)
	}
	Mock2 = false

	// if err != nil {
	// 	panic(err)
	// }
}

func TestModeChoices(t *testing.T) {

	testServer := httptest.NewServer(server())
	defer testServer.Close()

	req := httptest.NewRequest("GET", "http://localhost:3000/modify/", nil)
	rec := httptest.NewRecorder()

	var i io.ReadSeeker
	i, _ = os.Open("tmp/input.jpeg")

	err := renderModeChoices(rec, req, i, "jpeg")
	if err != nil {
		panic(err)
	}

	Mock3 = true
	err = renderModeChoices(rec, req, i, "jpeg")
	if err != nil {
		panic(err)
	}
	Mock3 = false

	Mock4 = true
	err = renderModeChoices(rec, req, i, "jpeg")
	if err != nil {
		panic(err)
	}
	Mock4 = false

}

func TestGenImages(t *testing.T) {
	opts := []genOpts{
		{N: 10, M: primitive.ModeCircle},
	}
	var i io.ReadSeeker
	i, _ = os.Open("tmp/input.jpeg")

	_, err := genImages(i, "jpeg", opts...)
	if err != nil {
		panic(err)
	}

	Mock5 = true
	_, err = genImages(i, "jpeg", opts...)
	if err != nil {
		panic(err)
	}
	Mock5 = false

}

func TestGenImage(t *testing.T) {
	// var i io.Reader
	i1, _ := os.Open("tmp/input.jpeg")
	i2, _ := os.Open("tmp/input.jpeg")
	i3, _ := os.Open("tmp/input.jpeg")

	_, err := genImage(i1, "jpeg", 10, primitive.ModeBeziers)
	if err != nil {
		panic(err)
	}

	Mock6 = true
	_, err = genImage(i2, "jpeg", 10, primitive.ModeBeziers)
	if err != nil {
		panic(err)
	}
	Mock6 = false

	Mock7 = true
	_, err = genImage(i3, "jpeg", 10, primitive.ModeBeziers)
	if err != nil {
		panic(err)
	}
	Mock7 = false
}

func TestTempFile(t *testing.T) {

	_, err := Tempfile("", "")
	if err != nil {
		panic(err)
	}

	Mock8 = true
	_, err = Tempfile("", "")
	if err != nil {
		panic(err)
	}
	Mock8 = false

}

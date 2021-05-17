package main

import (
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"testing"
)

func TestMain(t *testing.T) {
	temp1 = true
	main()
}

func TestMakeLinks(t *testing.T) {
	stack := debug.Stack()
	link := makeLinks(string(stack))
	if link == "" {
		t.Error("Link is Empty")
	}
}

func TestSourceHandler(t *testing.T) {
	testCaseTable := []struct {
		testCaseName string
		link         string
		statusCode   int
	}{
		{
			testCaseName: "test1",
			link:         "line=72&path=/home/gslab/go/src/github.com/abhishek-devani/Gophercises/recover_middleware_w_source_code/main.go",
			statusCode:   200,
		},
		{
			testCaseName: "test2",
			link:         "line=24&path=/usr/local/go/src/runtime/debug/stack.go",
			statusCode:   200,
		},
	}

	for i := 0; i < len(testCaseTable); i++ {
		// var err error
		req, err := http.NewRequest("GET", "http://localhost:3000/debug/?"+testCaseTable[i].link, nil)
		if err != nil {
			t.Fatalf("not able to run")
		}
		rec := httptest.NewRecorder()
		Mock1 = true
		sourceCodeHandler(rec, req)
		Mock1 = false
		Mock2 = true
		sourceCodeHandler(rec, req)
		Mock2 = false
		Mock3 = true
		sourceCodeHandler(rec, req)
		Mock3 = false
		Mock4 = true
		sourceCodeHandler(rec, req)
		Mock4 = false
		sourceCodeHandler(rec, req)

		res := rec.Result()
		if res.StatusCode != testCaseTable[i].statusCode {
			t.Errorf("Test case Number: %v Expected %v , Actual status %v", testCaseTable[i].testCaseName, testCaseTable[i].statusCode, res.StatusCode)
		}
	}
}

func TestMiddleware(t *testing.T) {
	temp = true
	mux := http.HandlerFunc(panicDemo)

	req, err := http.NewRequest("Get", "/panic/", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	devMw(mux).ServeHTTP(rr, req)
}

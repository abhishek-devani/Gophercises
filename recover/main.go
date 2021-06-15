package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var temp bool
var Mock1 bool
var Mock2 bool
var Mock3 bool
var Mock4 bool
var Mock5 bool

var temp1 bool

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/debug/", sourceCodeHandler)
	mux.HandleFunc("/panic/", panicDemo)
	if temp1 {
		server := &http.Server{Addr: ":3000", Handler: devMw(mux)}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	} else {
		log.Fatal(http.ListenAndServe(":3000", devMw(mux)))
	}
}

// It handles the source code of given file
func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, err := strconv.Atoi(lineStr)
	if err != nil || Mock1 {
		return
	}
	file, err := os.Open(path) // For read access.
	if err != nil || Mock2 {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, file)
	if err != nil || Mock3 {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}

	// lexer is a software component that analyzes a string and breaks it up into its component parts
	lexer := lexers.Get("go")
	iterator, _ := lexer.Tokenise(nil, b.String())
	style := styles.Get("github")
	if style == nil || Mock4 {
		return
		// style = styles.Fallback
	}
	formatter := html.New(html.TabWidth(2), html.HighlightLines(lines))
	w.Header().Set("content-type", "text/html")
	fmt.Fprint(w, "<style>pre { font-size: 1.2em; }</style>")
	formatter.Format(w, style, iterator)
}

// Middleware to recover the panics
func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil || Mock5 {
				log.Println(err)
				stack := debug.Stack()
				// log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

// Handle Panic Page
func panicDemo(w http.ResponseWriter, r *http.Request) {
	if !temp {
		panic("Oh no!")
	}
}

// convert the stack into Links
func makeLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		file := ""
		for i, ch := range line {
			if ch == ':' {
				file = line[1:i]
				// fmt.Println(file)
				break
			}
		}
		var lineStr strings.Builder
		for i := len(file) + 2; i < len(line); i++ {
			// fmt.Printf("len(file):%v, len(line):%v\n", len(file)+2, len(line))
			if line[i] < '0' || line[i] > '9' {
				break
			}
			// fmt.Println(line[i])
			lineStr.WriteByte(line[i])
		}
		v := url.Values{}
		v.Set("path", file)
		v.Set("line", lineStr.String())
		lines[li] = "\t<a href=\"/debug/?" + v.Encode() + "\">" + file + ":" + lineStr.String() + "</a>" +
			line[len(file)+2+len(lineStr.String()):]
	}
	return strings.Join(lines, "\n")
}

// Intital Page
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

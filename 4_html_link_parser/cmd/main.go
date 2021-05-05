package main

import (
	"fmt"
	"strings"

	link "github.com/abhishek-devani/Gophercises/4_html_link_parser"
)

var exampleHtml = `
    <html>
    <body>
      <h1>Hello!</h1>
      <a href="/other-page">A link to another page</a>
      <a href="/page-two">
        A link to second page
        <span> some span </span>
      </a>
    </body>
    </html>
		`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

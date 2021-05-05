package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	link "github.com/abhishek-devani/Gophercises/4_html_link_parser"
)

/*
	1. GET the web page.
	2. parse all the links on the page
	3. build proper urls with our links
	4. filter out any links with different domain
	5. find all the pages (BFS)
	6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com​", "the url that you want ot build sitemap")
	flag.Parse()

	// https://demo.qodeinteractive.com/bridge32/

	fmt.Println(*urlFlag)

	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}

}

func get(urlStr string) []string {

	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)

	// link after all redirection
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(base, hrefs(resp.Body, base))
}

func hrefs(body io.Reader, base string) []string {
	links, _ := link.Parse(body)
	var ret []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(base string, links []string) []string {
	var ret []string
	for _, link := range links {
		if strings.HasPrefix(link, base) {
			ret = append(ret, link)
		}
	}
	return ret
}

package main

import (
	"testing"
)

func TestCode(t *testing.T) {
	inp := "https://demo.qodeinteractive.com/bridge32/"

	pages := bfs(inp, 1)
	XmlEncode(pages)

}
